package ansi

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type Writer struct {
	out io.Writer
}

func NewAnsiStdout() io.Writer {
	return &Writer{
		out: os.Stdout,
	}
}

func NewAnsiStderr() io.Writer {
	return &Writer{
		out: os.Stderr,
	}
}

func (w *Writer) Write(data []byte) (n int, err error) {
	r := bytes.NewReader(data)

	for {
		ch, size, err := r.ReadRune()
		if err != nil {
			break
		}
		n += size

		switch ch {
		case '\\':
			w.trapBashSlash(r)
		default:
			fmt.Print(string(ch))
		}
	}
	return
}

func (w *Writer) trapBashSlash(r *bytes.Reader) {
}
