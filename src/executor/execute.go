// Package executor executes commands that are given to it.
package executor

import (
	"fmt"
	"os"
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
    // Pipe output & input to terminal
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Run()
  }
  return nil
}

// printError prints the error with the information given.
// procname and command are printed to tell the user which command
// returned an error code, and err is the information from stderr
// regarding the error.
func printError(procname string, command []string, err string) {
  fmt.Println("gsp failed at the following command:")
  // Make the command into a better-looking string
  s := procname + " "
  for i := 0; i < len(command); i++ {
    s += command[i] + " "
  }
  fmt.Println(s)
  fmt.Println("Received message from stderr:")
  fmt.Println(err)
}

// commandToString converts command to a string, where each element
// in command is separated by a whitespace (since terminal options are
// separated by whitespace). The first word in the value returned by 
// comandToString is procname.
func commandToString(procname string, command []string) string {
  s := procname + " "
  len := len(command)
  for i := 0; i < len - 1; i++ {
    s += command[i] + " "
  }
  s += command[len-1]
  return s
}
