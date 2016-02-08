shell [![Build Status](https://travis-ci.org/keegancsmith/shell.svg?branch=master)](https://travis-ci.org/) [![GoDoc](https://godoc.org/github.com/keegancsmith/shell?status.svg)](https://godoc.org/github.com/keegancsmith/shell)
======

Generate Shell Commands in Go, sprintf Style. Inspired by libphutil.

Quick example:

```go
// Generates shell command to find number of go files on a remote machine
host := "foo.com"
findCmd := shell.Sprintf("find . -iname %s | wc -l", "*.go")
remoteCmd := shell.Sprintf("ssh ubuntu@%s %s", host, findCmd)
fmt.Println(remoteCmd)
// Output: ssh ubuntu@foo.com 'find . -iname '\''*.go'\'' | wc -l'
```

Also slightly extends `Sprintf` syntax to allow easily passing in slices. This
is a common use case when interacting with commands that take in a list of
arguments:

```go
// Support for passing in a slice of arguments
gitcmd := shell.Sprintf("git add %S", []string{"foo.go", "bar.go", "test data"})
fmt.Println(gitcmd)
// Output: git add foo.go bar.go 'test data'
```

See https://godoc.org/github.com/keegancsmith/shell for more information.
