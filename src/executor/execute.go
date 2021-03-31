// Package executor executes commands that are given to it.
package executor

import (
	"os"
	"os/exec"

	errhandler "github.com/rohit-px2/gitsplit/src/errors"
)

// Execute runs each command in commands as a system command
// using exec.Command, for a given
// process. If all commands are run successfully,
// the function returns nil. Otherwise, the function will produce
// an error.
// Execute also prints the output of each command to stdout.
// Requirements:
// procname must be callable from the command line.
// If a command produces an error then it will be returned from Execute.
func Execute(procname string, commands [][]string) {
	for _, command := range commands {
		cmd := exec.Command(procname, command...)
		// Pipe output & input to terminal
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		// Terminate if git gives us an error.
		errhandler.CheckExitFatal(err)
	}
}

// IsExecutable returns 'true' if a command-line command can be executed.
func IsExecutable(procname string) bool {
	_, err := exec.LookPath(procname)
	return err == nil
}
