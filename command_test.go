package shell

import "testing"

func TestCommandf(t *testing.T) {
	out, err := Commandf("echo %s  world   %s", "hello", "quote!me").Output()
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "hello world quote!me\n" {
		t.Fatalf("Unexpected hello world output: %#v", string(out))
	}
}
