// Copyright 2017 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/maruel/subcommands"

	"go.chromium.org/luci/common/errors"
)

var lockFileName = "mmutex.lock"
var lockFileEnvVariable = "MMUTEX_LOCK_DIR"
var fslockTimeout = 2 * time.Hour
var fslockPollingInterval = 5 * time.Second

// Returns the lock file path based on the environment variable, or an empty string if no
// lock file should be used.
func computeLockFilePath(env subcommands.Env) (string, error) {
	envVar := env[lockFileEnvVariable]
	if !envVar.Exists {
		return "", nil
	}

	lockFileDir := envVar.Value
	if !filepath.IsAbs(lockFileDir) {
		return "", errors.Reason("Lock file directory %s must be an absolute path", lockFileDir).Err()
	}

	if _, err := os.Stat(lockFileDir); os.IsNotExist(err) {
		fmt.Printf("Lock file directory %s does not exist, mmutex acting as a passthrough.", lockFileDir)
		return "", nil
	}

	return filepath.Join(lockFileDir, lockFileName), nil
}

var application = &subcommands.DefaultApplication{
	Name: "mmutex",
	Title: `'Maintenance Mutex' - Global mutex to isolate maintenance tasks.

mmutex is a command line tool that helps prevent maintenance tasks from running
during user tasks. The tool does this by way of a global lock file that users
must acquire before running their tasks.

Clients can use this tool to request that their task be run with one of two
types of access to the system:

  * Exclusive access guarantees that no other callers have any access
    exclusive or shared) to the resource while the specified command is run.
  * Shared access guarantees that only other callers with shared access
    will have access to the resource while the specified command is run.

In short, exclusive access guarantees a task is run alone, while shared access
tasks may be run alongside other shared access tasks.

The source for mmutex lives at:
  https://github.com/luci/luci-go/tree/master/mmutex`,
	Commands: []*subcommands.Command{
		cmdExclusive,
		cmdShared,
		subcommands.CmdHelp,
	},
	EnvVars: map[string]subcommands.EnvVarDefinition{
		"MMUTEX_LOCK_DIR": {
			ShortDesc: "The directory containing the lock and drain files.",
			Default:   "",
		},
	},
}

func main() {
	os.Exit(subcommands.Run(application, nil))
}
