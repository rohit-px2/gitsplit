// Package split ...
// This package contains a function / series of functions
// to split command-line arguments by a list of specified strings.
// Since this is within a git-command context, we need to keep track
// of context.
// For example, "add" is a git command, but "remote add" is also a git command
// which does something else, so we need to know when the user specifies "remote"
// before "add".
package split

import "errors"
// SplitByCommands splits args according to the
// given commands.
// Requirements:
// len(args) > 0
// If len(commands) = 0 then we would return args (nothing to split by)
// Here is a description of the algorithm:
// 1. For each argument, look for a command that matches the argument
// 2. If we find a command, we check whether there is a multi-word command
// that fits the command better (ex. "remote add" vs "add").
// 3. If the argument is not a command, add it to the current command we are
// processing (it's part of a command).
func SplitByCommands(args []string, commands []string) ([][]string, error) {
  numArgs := len(args)
  if numArgs <= 0 {
    return nil, errors.New("args must have at least one item")
  }
  totalSplit := make([][]string, 0)
  currentSplit := make([]string, 0)
  
  for i := 0; i < numArgs; i++ {
    // Check if we have a longer match (e.g. "remote add")
    // "remote" would be in commands[0]
    // We check if currentSplit's length is greater than 0
    // and if currentSplit[0] + " " + args[i] is a command.
    if len(currentSplit) > 0 && containsString(commands, currentSplit[0] + " " + args[i]) {
      currentSplit = append(currentSplit, args[i])
    } else if containsString(commands, args[i]) {
      totalSplit = append(totalSplit, currentSplit)
      currentSplit = nil
      currentSplit = append(currentSplit, args[i])
    } else {
      currentSplit = append(currentSplit, args[i])
    }

    if i == numArgs-1 {
      totalSplit = append(totalSplit, currentSplit)
      currentSplit = nil
    }
  }
  return totalSplit, nil
}

// containsString(arr, s) returns 'true' if s is an element of arr.
func containsString(arr []string, s string) bool {
  len := len(arr)
  for i := 0; i < len; i++ {
    if arr[i] == s {
      return true
    }
  }
  return false
}
