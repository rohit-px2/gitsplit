// Package constants holds the constant
// values for gitsplit, such as lists of Git commands,
// and possibly more.
// No other part of the code should access these variables
// directly, only using the functions to access them,
// so that there is no mutation.
package constants

var commands = []string{
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

// GetGitCommands returns a list of all Git commands.
func GetGitCommands() []string {
	return commands
}
