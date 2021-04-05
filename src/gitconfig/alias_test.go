package gitconfig

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpand(t *testing.T) {
	config := map[string][]string{
		"unstage": {"reset", "HEAD", "--"},
		"a":       {"add"},
		"com":     {"commit"},
	}
	list := []string{
		"a", ".", "com", "push", "origin", "main",
	}
	expect := []string{
		"add", ".", "commit", "push", "origin", "main",
	}
	result := Expand(list, config)
	assert.Equal(t, expect, result)
}

func TestIsAlias(t *testing.T) {
	aliases := map[string][]string{
		"unstage": {"reset", "HEAD", "--"},
		"com":     {"commit"},
		"a":       {"add"},
	}
	potentialAlias := "a"
	isAlias, expanded := IsAlias(aliases, potentialAlias)
	assert.True(t, isAlias)
	expectExpanded := []string{"add"}
	assert.Equal(t, expectExpanded, expanded)
	potentialAlias = "b"
	isAlias, expanded = IsAlias(aliases, potentialAlias)
	assert.False(t, isAlias)
	expectExpanded = []string{}
	assert.Equal(t, expectExpanded, expanded)
}
