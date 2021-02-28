package main

import (
	"fmt"
	"log"
	"os"
	"github.com/rohit-px2/gitsplit/src/split"
)

// A list of every Git command
var commands = []string {
  "config",
  "help",
  "bugreport",
  "init",
  "clone",
  "add",
  "status",
  "diff",
  "commit",
  "notes",
  "restore",
  "reset",
  "rm",
  "mv",
  "branch",
  "checkout",
  "switch",
  "merge",
  "mergetool",
  "log",
  "stash",
  "tag",
  "worktree",
  "fetch",
  "pull",
  "push",
  "remote",
  "remote add",
  "remote rename",
  "remote get-url",
  "remote set-url", 
  "remote set-head",
  "remote remove",
  "remote set-branches",
  "remote prune",
  "submodule",
  "show",
  "log",
  "diff",
  "difftool",
  "range-diff",
  "shortlog",
  "describe",
  "apply",
  "cherry-pick",
  "rebase",
  "revert",
  "bisect",
  "blame",
  "grep",
  "am",
  "apply",
  "format-patch",
  "send-email",
  "request-pull",
  "svn",
  "fast-import",
  "clean",
  "gc",
  "fsck",
  "reflog",
  "filter-branch",
  "instaweb",
  "archive",
  "bundle",
  "daemon",
  "update-server-info",
  "cat-file",
  "check-ignore",
  "checkout-index",
  "commit-tree",
  "count-objects",
  "diff-index",
  "for-each-ref",
  "hash-object",
  "ls-files",
  "ls-tree",
  "merge-base",
  "read-tree",
  "rev-list",
  "rev-parse",
  "show-ref",
  "symbolic-ref",
  "update-index",
  "update-ref",
  "verify-pack",
  "write-tree",
}

func main() {
  args := os.Args[1:]
  // Display help message if there are no arguments
  if len(args) < 1 {
    displayHelp()
    os.Exit(0)
  }
  splits, err := split.SplitByCommands(args, commands)
  if err != nil {
    log.Fatal(err)
  }
  // Commands are split into git commands.
  // We can chain the commands by adding "git" before each of them
  // and adding "&&" to the end of them (except for the last command)
  res := fmt.Sprintf("Length: %d\n", len(splits))
  fmt.Println(res)
  for i := 0; i < len(splits); i++ {
    fmt.Println(splits[i])
  }
}

func displayHelp() {
  fmt.Println("Gitsplit v1.0")
  fmt.Println("Just run git commands (add, init, ...)")
}
