package executor

import (
	"fmt"
	"strings"
)

// printError prints the error with the information given.
// procname and command are printed to tell the user which command
// returned an error code, and err is the information from stderr
// regarding the error.
func printError(procname string, command []string, err string) {
	fmt.Println("gsp failed at the following command:")
	// Make the command into a better-looking string
	s := procname + " "
	for _, elem := range command {
		s += elem + " "
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
	var sb strings.Builder
	sb.WriteString(procname)
	sb.WriteString(" ")
	for _, elem := range command {
		sb.WriteString(elem)
		sb.WriteString(" ")
	}
	sb.WriteString(command[len(command)-1])
	return sb.String()
}
