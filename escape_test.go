package shell

import "testing"

func TestEscapeArg(t *testing.T) {
	cases := []struct{ in, want string }{
		{`Mr O'Neil`, `'Mr O'\''Neil'`},
		{`Mr O\'Neil`, `'Mr O\'\''Neil'`},
		{`%FILENAME`, `'%FILENAME'`},
		{``, `''`},
	}
	for _, c := range cases {
		got := EscapeArg(c.in)
		if got != c.want {
			t.Fatalf("got %v, wanted %v for %v", got, c.want, c.in)
		}
	}
}
