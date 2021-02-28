// Package executor executes commands that are given to it.
package executor

import (
	"fmt"
	"os/exec"
)

// Execute runs each command in commands as a system command
// using exec.Command, for a given
// process. If all commands are run successfully,
// the function returns nil. Otherwise, the function will produce
// an error.
// Execute also prints the output of each command to stdout.
// Requirements:
// procname musut be callable from the command line.
// If a command produces an error then it will be returned from Execute.
func Execute(procname string, commands [][]string) (error) {
  for _, command := range commands {
    cmd := exec.Command(procname, command...)
    res, err := cmd.Output()
    if err != nil {
      return err
    }
    fmt.Println(string(res))
  }
  return nil
}