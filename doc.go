// Package shell provides utilities for building command strings to execute in
// sh. It also provides a wrapper around os/exec, but instead of a Args slice
// you can pass in a format string.
//
// A common use case is generating safe strings for passing to commands that
// execute other shell commands (such as SSH). Otherwise you may find this a
// more convenient way to build up commands vs passing a slice to exec.Command
//
// This way of interacting with the shell is inspired by the csprintf related
// functions from libphutil
package shell
