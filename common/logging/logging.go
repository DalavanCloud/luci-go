// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

/*
Package logging defines Logger interface and context.Context helpers to put\get
logger from context.Context.

Unfortunately standard library doesn't define any Logger interface (only
struct). And even worse: GAE logger is exposing different set of methods. Some
additional layer is needed to unify the logging. Package logging is intended to
be used from packages that support both local and GAE environments. Such
packages should not use global logger but must accept instances of Logger
interface (or even more generally context.Context) as parameters. Then callers
can pass appropriate Logger implementation  (or inject appropriate logger into
context.Context) depending on where the code is running.

Libraries under luci-go/common/ MUST use luci-go/common/logging instead of
directly instantiating concrete implementations.
*/
package logging

import (
	"golang.org/x/net/context"
)

// Logger interface is ultimately implemented by underlying logging libraries
// (like go-logging or GAE logging). It is the least common denominator among
// logger implementations.
type Logger interface {
	// Debugf formats its arguments according to the format, analogous to
	// fmt.Printf and records the text as a log message at Debug level.
	Debugf(format string, args ...interface{})

	// Infof is like Debugf, but logs at Info level.
	Infof(format string, args ...interface{})

	// Warningf is like Debugf, but logs at Warning level.
	Warningf(format string, args ...interface{})

	// Errorf is like Debugf, but logs at Error level.
	Errorf(format string, args ...interface{})

	// LogCall is a generic logging function. This is oriented more towards
	// utility functions than direct end-user usage.
	LogCall(l Level, calldepth int, format string, args []interface{})
}

// Factory is a method that returns a Logger instance for the specified context.
type Factory func(context.Context) Logger

type key int

const (
	loggerKey key = iota
	fieldsKey
	levelKey
)

// SetFactory sets the Logger factory for this context.
//
// The factory will be called each time Get(context) is used.
func SetFactory(c context.Context, f Factory) context.Context {
	return context.WithValue(c, loggerKey, f)
}

// Set sets the logger for this context.
//
// It can be retrieved with Get(context).
func Set(c context.Context, l Logger) context.Context {
	return SetFactory(c, func(context.Context) Logger { return l })
}

// GetFactory returns the currently-configured logging factory.
func GetFactory(c context.Context) Factory {
	if f, ok := c.Value(loggerKey).(Factory); ok {
		return f
	}
	return NullFactory
}

// Get the current Logger, or a logger that ignores all messages if none
// is defined.
func Get(c context.Context) (ret Logger) {
	return GetFactory(c)(c)
}
