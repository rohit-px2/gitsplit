package main

import (
	"fmt"
	"github.com/rohit-px2/gitsplit/src/constants"
  "github.com/rohit-px2/gitsplit/src/errors"
  "github.com/rohit-px2/gitsplit/src/executor"
	"github.com/rohit-px2/gitsplit/src/split"
	"os"
)

func main() {
	args := os.Args[1:]
	// Display help message if there are no arguments
	if len(args) < 1 {
		displayHelp()
		os.Exit(0)
	}
	commands := constants.GetGitCommands()
	splits, err := split.SplitByCommands(args, commands)
  errors.CheckFatal(err)
	// Commands are split into git commands.
	// We can chain the commands by adding "git" before each of them
	progName := "git"
  if !executor.IsExecutable(progName) {
    displayInstallGitMessage()
    os.Exit(0)
  }
  err = executor.Execute(progName, splits)
  errors.CheckFatal(err)
}

// displays the Help message, showing the version and how to use the program.
func displayHelp() {
  msg := 
  `
GitSplit v1.1
Usage: gsp [commands]
Commands are git commands and are executed through git.
Inputting multiple commands acts the same as if you were to put "&&" between
each command.

Example:
gsp init add . commit -m "initial commit"
would result in the same effect as
git init && git add . && git commit -m "initial commit".

For information on git commands, you can find the documentation at
https://git-scm.com/docs.
  `
  fmt.Print(msg)
}

// displayInstallGitMessage does what the name says it does
func displayInstallGitMessage() {
  msg :=
  `
It appears that you do not have Git installed. Git is needed to run GitSplit.
You can download git from git-scm.com.
  `
  fmt.Print(msg)
}
