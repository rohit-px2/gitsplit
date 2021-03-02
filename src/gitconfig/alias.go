// Package gitconfig contains functions
// for getting gitconfig data from local, global, and system
package gitconfig

import (
	"bytes"
	"os/exec"
	"strings"
)

// GetAllMatching gets all git config variables that match the given regex.
func GetAllMatching(regex string) []string {
	var stdout bytes.Buffer
	cmd := exec.Command("git", "config", "--get-regexp", regex)
	cmd.Stdout = &stdout
	err := cmd.Run()
	// If we get an error we go without the config
	if err != nil {
		return []string{}
	}
	out := stdout.String()
	res := strings.Split(out, " ")
	return res
}
