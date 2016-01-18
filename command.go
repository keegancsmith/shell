package shell

import (
	"fmt"
	"io"
	"os/exec"
)

// Commandf runs a shell command based on the format string
func Commandf(format string, a ...interface{}) *exec.Cmd {
	shellCmd := Sprintf(format, a...)
	return exec.Command("/bin/sh", "-c", shellCmd)
}

// Sprintf generates a shell command with the format arguments escaped.
func Sprintf(format string, a ...interface{}) string {
	wrapped := make([]interface{}, len(a))
	for i, v := range a {
		wrapped[i] = escapable{v}
	}
	return fmt.Sprintf(format, wrapped...)
}

type escapable struct {
	x interface{}
}

func (e escapable) Format(f fmt.State, c rune) {
	s := "%"
	for i := 0; i < 128; i++ {
		if f.Flag(i) {
			s += string(i)
		}
	}
	if w, ok := f.Width(); ok {
		s += fmt.Sprintf("%d", w)
	}
	if p, ok := f.Precision(); ok {
		s += fmt.Sprintf(".%d", p)
	}
	s += string(c)
	formatted := fmt.Sprintf(s, e.x)
	io.WriteString(f, ReadableEscapeArg(formatted))
}
