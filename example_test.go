package shell_test

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/keegancsmith/shell"
)

func ExampleSprintf() {
	// Generates shell command to find number of go files on a remote machine
	host := "foo.com"
	findCmd := shell.Sprintf("find . -iname %s | wc -l", "*.go")
	remoteCmd := shell.Sprintf("ssh ubuntu@%s %s", host, findCmd)
	fmt.Println(remoteCmd)
	// Output: ssh ubuntu@foo.com 'find . -iname '\''*.go'\'' | wc -l'
}

func ExampleCommandf() {
	out, err := shell.Commandf("echo %s", "hello world").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(out))
	// Output: hello world
}

func ExampleCommandf_redirect() {
	var stdout, stderr bytes.Buffer
	cmd := shell.Commandf("echo %s; echo %s 1>&2", "hello from stdout", "hello from stderr")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("stdout:", strings.TrimSpace(stdout.String()))
	fmt.Println("stderr:", strings.TrimSpace(stderr.String()))
	// Output: stdout: hello from stdout
	// stderr: hello from stderr
}
