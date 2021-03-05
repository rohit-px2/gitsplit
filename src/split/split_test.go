package split

import (
	"github.com/rohit-px2/gitsplit/src/constants"
	"github.com/stretchr/testify/assert"
	"testing"
  "strings"
)

// Testing git commands and seeing whether SplitByCommands
// splits them appropriately
func TestSplitByCommands(t *testing.T) {
  // Test 1
  emptyconf := map[string] []string{}
	commands := constants.GetGitCommands()
	arguments := []string{
		"init",
		"add",
		".",
		"remote",
		"add",
		"origin",
		"https://github.com/rohit-px2/gitsplit.git",
	}
	expect := [][]string{
		{"init"},
		{"add", "."},
		{"remote", "add", "origin", "https://github.com/rohit-px2/gitsplit.git"},
	}
	splits, err := SplitByCommands(arguments, commands, emptyconf)
	assert.Nil(t, err)
	assert.Equal(t, expect, splits, "they should be equal")
  // Test 2
  command := "add . commit -m \"Hello World\" push origin main"
  arguments = strings.Split(command,  " ")
  expect = [][]string {
    {"add", "."},
    {"commit", "-m", "\"Hello", "World\""},
    {"push", "origin", "main"},
  }

  splits, err = SplitByCommands(arguments, commands, emptyconf)
  assert.Nil(t, err)
  assert.Equal(t, expect, splits, "commit comments should be an escaped string")

  config := map[string] []string {
    "unstage": {"reset", "HEAD", "--"},
    "com": {"commit"},
  }
  command = "add . com -m \"testing.. \" unstage file.txt"
  arguments = strings.Split(command, " ")
  expect = [][]string {
    {"add", "."},
    {"com", "-m", "\"testing...", "\""},
    {"unstage", "file.txt"},
  }
  splits, err = SplitByCommands(arguments, commands, config)
  assert.Nil(t, err)
  assert.Equal(t, expect, splits, "config vars should be recognized")

}

func TestAdding(t *testing.T) {
  list := []string {"0", "1", "2", "3"}
  index := 2
  valToAdd := "added"
  expect := []string {"0", "1", "added", "2", "3"}
  result := add(list, valToAdd, index)
  assert.Equal(t, expect, result, `add should shift values to the right`)
}
