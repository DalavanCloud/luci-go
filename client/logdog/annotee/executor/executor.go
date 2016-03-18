// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package executor

import (
	"errors"
	"fmt"
	"io"
	"os/exec"
	"syscall"

	"github.com/luci/luci-go/client/logdog/annotee"
	"github.com/luci/luci-go/client/logdog/annotee/annotation"
	"github.com/luci/luci-go/common/ctxcmd"
	"github.com/luci/luci-go/common/logdog/types"
	"github.com/luci/luci-go/common/proto/milo"
	"golang.org/x/net/context"
)

// AnnotationMode describes how the Executor will process annotations.
type AnnotationMode int

const (
	// NoAnnotations causes no annotation processing will be performed on the
	// bootstrapped process' STDOUT.
	NoAnnotations AnnotationMode = iota
	// TeeAnnotations causes the bootstrapped process' annotation state to be
	// transmitted through LogDog as an annotation stream, but still included in
	// the bootstrapped process' STDOUT stream.
	TeeAnnotations
	// StripAnnotations causes the bootstrapped process' annotation state to be
	// transmitted through LogDog as an annotation stream and removed from the
	// bootstrapped process' STDOUT stream.
	StripAnnotations
)

// Executor bootstraps an application, running its output through a Processor.
type Executor struct {
	// Options are the set of Annotee options to use.
	Options annotee.Options

	// Annoate describes how annotations in the STDOUT stream should be handled.
	Annotate AnnotationMode

	// Stdin, if not nil, will be used as standard input for the bootstrapped
	// process.
	Stdin io.Reader

	// TeeStdout, if not nil, is a Writer where bootstrapped process standard
	// output will be tee'd.
	TeeStdout io.Writer
	// TeeStderr, if not nil, is a Writer where bootstrapped process standard
	// error will be tee'd.
	TeeStderr io.Writer

	returnCode int
	steps      []*milo.Step
}

// Run executes the bootstrapped process, blocking until it completes.
func (e *Executor) Run(ctx context.Context, command []string) error {
	// Clear any previous state.
	e.returnCode = 0
	e.steps = nil

	if len(command) == 0 {
		return errors.New("no command")
	}

	ctx, cancelFunc := context.WithCancel(ctx)
	cmd := ctxcmd.CtxCmd{
		Cmd: exec.Command(command[0], command[1:]...),
	}

	// STDOUT
	stdoutRC, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to create STDOUT pipe: %s", err)
	}
	defer stdoutRC.Close()
	stdout := e.configStream(stdoutRC, types.StreamName("stdout"), e.TeeStdout)

	stderrRC, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("failed to create STDERR pipe: %s", err)
	}
	defer stderrRC.Close()
	stderr := e.configStream(stderrRC, types.StreamName("stderr"), e.TeeStderr)

	// Start our process.
	if err := cmd.Start(ctx); err != nil {
		return fmt.Errorf("failed to start bootstrapped process: %s", err)
	}
	processRunning := true
	defer func() {
		cancelFunc()
		if processRunning {
			_ = cmd.Wait()
		}
	}()

	// Infer our execution.
	options := e.Options
	if options.Execution == nil {
		options.Execution = annotation.ProbeExecution(command)
	}

	// Configure our Processor.
	streams := []*annotee.Stream{
		stdout,
		stderr,
	}

	proc := annotee.New(ctx, options)
	defer proc.Finish()

	if err := proc.RunStreams(streams); err != nil {
		return fmt.Errorf("failed to run processor: %s", err)
	}

	// Wait for our command to finish.
	if err := cmd.Wait(); err != nil {
		switch err.(type) {
		case *exec.ExitError:
			status := err.(*exec.ExitError).Sys().(syscall.WaitStatus)
			e.returnCode = status.ExitStatus()

		default:
			return fmt.Errorf("failed to wait for bootstrapped process: %s", err)
		}
	} else {
		e.returnCode = 0
	}
	processRunning = false

	// Record our annotation steps.
	proc.Finish().ForEachStep(func(s *annotation.Step) {
		e.steps = append(e.steps, s.Proto())
	})

	return nil
}

// Steps returns a list of Steps from the latest run.
func (e *Executor) Steps() []*milo.Step {
	return e.steps
}

// ReturnCode returns the Executor's return code, or 0 if no execution has
// occurred.
func (e *Executor) ReturnCode() int {
	return e.returnCode
}

func (e *Executor) configStream(r io.Reader, name types.StreamName, tee io.Writer) *annotee.Stream {
	s := &annotee.Stream{
		Reader:           r,
		Name:             name,
		Tee:              tee,
		Alias:            "stdio",
		StripAnnotations: (e.Annotate == StripAnnotations),
	}
	if e.Annotate != NoAnnotations {
		s.Annotate = true
	}
	return s
}
