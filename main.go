package main

import (
	"fmt"
	"github.com/rohit-px2/gitsplit/src/constants"
	"github.com/rohit-px2/gitsplit/src/executor"
	"github.com/rohit-px2/gitsplit/src/split"
	"log"
	"os"
)

// A list of every Git command (hopefully)

func main() {
	args := os.Args[1:]
	// Display help message if there are no arguments
	if len(args) < 1 {
		displayHelp()
		os.Exit(0)
	}
	commands := constants.GetGitCommands()
	splits, err := split.SplitByCommands(args, commands)
	if err != nil {
		log.Fatal(err)
	}
	// Commands are split into git commands.
	// We can chain the commands by adding "git" before each of them
	progName := "git"
	err = executor.Execute(progName, splits)
	if err != nil {
		log.Fatal(err)
	}
}

// displays the Help message, showing the version and how to use the program.
func displayHelp() {
	fmt.Println("Gitsplit v1.0")
	fmt.Println("Just run git commands (add, init, ...)")
}
