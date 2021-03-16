// Package split contains a function / series of functions
// to split command-line arguments by a list of specified strings.
// Since this is within a git-command context, we need to keep track
// of context.
// For example, "add" is a git command, but "remote add" is also a git command
// which does something else, so we need to know when the user specifies "remote"
// before "add".
package split

import (
	"errors"

	"github.com/rohit-px2/gitsplit/src/gitconfig"
)

// ByCommands splits args according to the
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
func ByCommands(
  args      []string,
  commands  []string,
  config    map[string] []string,
) ([][]string, error) {
  args = gitconfig.Expand(args, config)
  numArgs := len(args)
  if numArgs <= 0 {
    return nil, errors.New("args must have at least one item")
  }
  totalSplit := make([][]string, 0)
  currentSplit := make([]string, 0)

	for i, elem := range args {
    // Check if we have a longer match (e.g. "remote add")
    // "remote" would be in commands[0]
    // We check if currentSplit's length is greater than 0
    // and if currentSplit[0] + " " + elem is a command.
    if len(currentSplit) > 0 && containsString(commands, currentSplit[0] + " " + elem) {
      currentSplit = append(currentSplit, elem)
    } else if containsString(commands, elem) {
      if i == 0 {
        currentSplit = append(currentSplit, elem)
      } else {
        totalSplit = append(totalSplit, currentSplit)
        currentSplit = nil
        currentSplit = append(currentSplit, elem)
      }
    } else {
      currentSplit = append(currentSplit, elem)
    }

    if i == numArgs-1 {
      totalSplit = append(totalSplit, currentSplit)
    }
  }
  return totalSplit, nil
}

// containsString(arr, s) returns 'true' if s is an element of arr.
func containsString(arr []string, s string) bool {
	for _, str := range arr {
		if str == s {
			return true
		}
	}
	return false
}

func add(arr []string, elem string, index int) ([]string, error) {
	size := len(arr)
	if index > size {
		return arr, errors.New("index out of bounds")
	} else if size == index {
    return append(arr, elem), nil
  }
  arr = append(arr[:index+1], arr[index:]...)
  arr[index] = elem
  return arr, nil
}

func removeIndex(arr []string, index int) []string{
  return append(arr[:index], arr[index+1:]...)
}
