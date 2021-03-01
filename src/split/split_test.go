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
	splits, err := SplitByCommands(arguments, commands)
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

  splits, err = SplitByCommands(arguments, commands)
  assert.Nil(t, err)
  assert.Equal(t, expect, splits, "commit comments should be an escaped string")
}
