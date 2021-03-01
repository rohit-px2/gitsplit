package executor

import (
	"fmt"
)

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
	for i := 0; i < len-1; i++ {
		s += command[i] + " "
	}
	s += command[len-1]
	return s
}
