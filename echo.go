// Package echo provides simple stdout output functions.
package echo

import "os"

// Necho writes s to stdout without a trailing newline.
func Necho(s string) {
	os.Stdout.Write([]byte(s))
}


// Echo writes s to stdout with a trailing newline.
func Echo(s string) {
	os.Stdout.Write([]byte(s + "\n"))
}