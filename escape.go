package shell

import (
	"regexp"
	"strings"
)

// EscapeArg escapes a string such that it is safe to pass to a shell.
// It is a re-implementation of PHP's escapeshellarg
func EscapeArg(arg string) string {
	// TODO support WIN32 and skipping non-valid multibyte characters
	return "'" + strings.Replace(arg, "'", `'\''`, -1) + "'"
}

// ReadableEscapeArg will not escape strings that do not requiring
// escaping. Note that it is conservative in it's approach, so may escape
// strings which do not require it.
func ReadableEscapeArg(arg string) string {
	if readableRe.MatchString(arg) {
		return arg
	}
	return EscapeArg(arg)
}

var readableRe = regexp.MustCompile("^[a-zA-Z0-9:/@._-]+$")
