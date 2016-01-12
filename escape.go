package shell

import "strings"

// EscapeArg escapes a string such that it is safe to pass to a shell.
// It is a re-implementation of PHP's escapeshellarg
func EscapeArg(arg string) string {
	// TODO support WIN32 and skipping non-valid multibyte characters
	return "'" + strings.Replace(arg, "'", `'\''`, -1) + "'"
}
