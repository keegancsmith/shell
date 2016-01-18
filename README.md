# shell

Utilities for interacting with the shell. Inspired by libphutil.

Quick example:

```go
// Generates shell command to find number of go files on a remote machine
host := "foo.com"
findCmd := shell.Sprintf("find . -iname %s | wc -l", "*.go")
remoteCmd := shell.Sprintf("ssh ubuntu@%s %s", host, findCmd)
fmt.Println(remoteCmd)
// Output: ssh ubuntu@foo.com 'find . -iname '\''*.go'\'' | wc -l'
```

See https://godoc.org/github.com/keegancsmith/shell for more information.
