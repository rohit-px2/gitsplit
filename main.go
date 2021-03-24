package main

import (
	"fmt"
	"os"

	"github.com/rohit-px2/gitsplit/src/constants"
	"github.com/rohit-px2/gitsplit/src/errors"
	"github.com/rohit-px2/gitsplit/src/executor"
	"github.com/rohit-px2/gitsplit/src/gitconfig"
	"github.com/rohit-px2/gitsplit/src/split"
)

func main() {
	args := os.Args[1:]
	// Display help message if there are no arguments
	parseOpts(args)
	conf := gitconfig.GetAliases()
	// If we have an error we continue with no alias variables
	commands := constants.GetGitCommands()
	splits, err := split.ByCommands(args, commands, conf)
	errhandler.CheckLogFatal(err)
	// Commands are split into git commands.
	// We can chain the commands by adding "git" before each of them
	progName := "git"
	if !executor.IsExecutable(progName) {
		displayInstallGitMessage()
		os.Exit(0)
	}
	executor.Execute(progName, splits)
}

// parseOpts determines which options the user specified and performs
// the corresponding action.
// NOTE: gsp should not override git flags
func parseOpts(args []string) {
	if len(args) < 1 {
		displayHelp()
	} else if len(args) == 1 && args[0] == "-v" {
		displayHelp()
	}
}

// displayHelp displays the Help message, showing the version and how to use the program.
func displayHelp() {
	msg :=
		`
GitSplit v1.1
Usage: gsp [-v] [<command> [<args>]]
<command> represents a git command, and is passed to git with <args>.
You can pass in multiple git commands and give arguments to each command.

Example:
gsp init add . commit -m "initial commit"
would result in the same effect as
git init && git add . && git commit -m "initial commit".

For information on git commands, you can find the documentation at
https://git-scm.com/docs.
  `
	fmt.Print(msg)
	os.Exit(0)
}

// displayInstallGitMessage displays a message to install Git when it is not detected
// on the user's computer
func displayInstallGitMessage() {
	msg :=
		`
ERROR:
It appears that you do not have Git installed. Git is needed to run GitSplit.
You can download git from https://git-scm.com.
  `
	fmt.Print(msg)
}
