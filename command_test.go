package shell

import (
	"bytes"
	"testing"
)

func TestCommandf(t *testing.T) {
	out, err := Commandf("echo %s  world   %s", "hello", "quote!me").Output()
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "hello world quote!me\n" {
		t.Fatalf("Unexpected hello world output: %#v", string(out))
	}
}

func TestCommandf_Redirect(t *testing.T) {
	var stdout, stderr bytes.Buffer
	cmd := Commandf("echo %s; echo %s 1>&2", "hello from stdout", "bye from stderr")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		t.Fatal(err)
	}
	if stdout.String() != "hello from stdout\n" {
		t.Errorf("Unexpected output from stdout: %v", stdout.String())
	}
	if stderr.String() != "bye from stderr\n" {
		t.Errorf("Unexpected output from stderr: %v", stderr.String())
	}
}
