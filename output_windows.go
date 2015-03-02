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
		case '\x1b':
			size, err = w.handleEscape(r)
			n += size
			if err != nil {
				break
			}
		default:
			fmt.Fprint(w.out, string(ch))
		}
	}
	return
}

func (w *Writer) handleEscape(r *bytes.Reader) (n int, err error) {
	buf := make([]byte, 0, 10)
	buf = append(buf, "\x1b"...)

	// Check '[' continues after \x1b
	ch, size, err := r.ReadRune()
	if err != nil {
		fmt.Fprint(w.out, string(buf))
		return
	}
	n += size
	if ch != '[' {
		fmt.Fprint(w.out, string(buf))
		return
	}

	// Parse escape code
	var code rune
	argBuf := make([]byte, 0, 10)
	for {
		ch, size, err = r.ReadRune()
		if err != nil {
			fmt.Fprint(w.out, string(buf))
			return
		}
		n += size
		if ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') {
			code = ch
			break
		}
		argBuf = append(argBuf, string(ch)...)
	}

	w.applyEscapeCode(buf, string(argBuf), code)
	return
}

func (w *Writer) applyEscapeCode(buf []byte, arg string, code rune) {
}
