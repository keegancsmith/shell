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

func TestReadableEscapeArg(t *testing.T) {
	cases := []struct {
		in       string
		readable bool
	}{
		{"helloworld", true},
		{"hello world", false},
		{"'", false},
		{"", false},
	}
	for _, c := range cases {
		var want string
		if c.readable {
			want = c.in
		} else {
			want = EscapeArg(c.in)
		}
		got := ReadableEscapeArg(c.in)
		if got != want {
			t.Fatalf("got %v, wanted %v for %v", got, want, c.in)
		}
	}
}
