package split

import (
	"github.com/rohit-px2/gitsplit/src/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplitByCommands(t *testing.T) {
	// copied from main.go file
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
}
