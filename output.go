// +build !windows

package ansi

import (
	"io"
	"os"
)

func NewAnsiStdout() io.Writer {
	return os.Stdout
}

func NewAnsiStderr() io.Writer {
	return os.Stderr
}
