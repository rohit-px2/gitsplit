// Package gitconfig contains functions
// for getting gitconfig data from local, global, and system
package gitconfig

import (
	"bytes"
	"os/exec"
	"strings"
)

// GetAliases returns a map of git alias variables to their values.
func GetAliases() map[string] []string {
	const aliasRegex = `^alias\.` // Regex match for "alias" config vars
  const aliasStart = "alias." // What to trim from the start of the config vars
	return GetAllMatching(aliasRegex, aliasStart)
}

// GetAllMatching gets all git config variables that match the given regex.
// These results are stored as an array of strings, where the element at each
// even-numbered index (0, 2, ...) is a configuration variable and the value
// to the index to the right of it is the configuration variable's value.
func GetAllMatching(regex string, trim string) map[string] []string {
	var stdout bytes.Buffer
	cmd := exec.Command("git", "config", "--get-regexp", regex)
  // output is captured from stdout
	cmd.Stdout = &stdout
	err := cmd.Run()
	// If we get an error we go without the config
	if err != nil {
    return map[string] []string {}
  }
  out := stdout.String()
  lines := strings.Split(out, "\n")
  // lines looks like this:
  // [
  //  "(config_var) (definition)"
  // ], where definition can be multiple words.
  // We want to change it into something like this:
  // {
  //  "(config_var)": "(definition)"
  // } to make it easier to access
  // we can split on the first newline we see to split each assignment
  // into the config_var (lhs) and the definition (rhs)
  // Then we turn this into a map by setting the key to lhs and the value to
  // rhs.
  m := make(map[string] []string)
  for _, line := range lines {
    // Split by first whitespace to get lhs & rhs
    firstWs := strings.Index(line, " ")
    if firstWs == -1 {
      continue
    }
    lhs := line[:firstWs]
    rhs := strings.Split(line[firstWs+1:], " ")
    // Right now, the rhs is one string, we turn that into an array of strings
    // split on whitespace so we don't have to split every time we use the map
    // Trim prefix of lhs (for example, we don't want "alias." prefix in
    // the map key)
    //lhs = strings.TrimPrefix(lhs, trim)
    lhs = strings.TrimPrefix(lhs, trim)
    m[lhs] = rhs
  }
  return m
}

// Expand expands the config variables in a list of strings to what their
// git aliases are.
func Expand(list []string, config map[string] []string) []string {
  current := make([]string, 0, 10)
	for _, elem := range list {
    if isAlias, expanded := IsAlias(config, elem); isAlias {
      current = append(current, expanded...)
    } else {
      current = append(current, elem)
    }
  }
  return current
}

// IsAlias determines whether str is a Git alias, and if it is, returns
// the string corresponding to the value of the Git alias corresponding to str.
// Also returns a boolean which depends on whether str is an alias or not.
func IsAlias(aliases map[string] []string, str string) (bool, []string) {
  if val, ok := aliases[str]; ok {
    return true, val
  }
	return false, []string{}
}
