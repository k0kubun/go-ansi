package ansi

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"syscall"
	"unsafe"

	"github.com/mattn/go-isatty"
)

var (
	kernel32                       = syscall.NewLazyDLL("kernel32.dll")
	procGetConsoleScreenBufferInfo = kernel32.NewProc("GetConsoleScreenBufferInfo")
	procSetConsoleTextAttribute    = kernel32.NewProc("SetConsoleTextAttribute")
)

type wchar uint16
type short int16
type dword uint32
type word uint16

type coord struct {
	x short
	y short
}

type smallRect struct {
	left   short
	top    short
	right  short
	bottom short
}

type consoleScreenBufferInfo struct {
	size              coord
	cursorPosition    coord
	attributes        word
	window            smallRect
	maximumWindowSize coord
}

type Writer struct {
	out     io.Writer
	handle  uintptr
	orgAttr uintptr
}

func NewAnsiStdout() io.Writer {
	var csbi consoleScreenBufferInfo
	out := os.Stdout
	if !isatty.IsTerminal(out.Fd()) {
		return out
	}
	handle := syscall.Handle(out.Fd())
	procGetConsoleScreenBufferInfo.Call(uintptr(handle), uintptr(unsafe.Pointer(&csbi)))
	return &Writer{out: out, handle: uintptr(handle), orgAttr: uintptr(csbi.attributes)}
}

func NewAnsiStderr() io.Writer {
	var csbi consoleScreenBufferInfo
	out := os.Stderr
	if !isatty.IsTerminal(out.Fd()) {
		return out
	}
	handle := syscall.Handle(out.Fd())
	procGetConsoleScreenBufferInfo.Call(uintptr(handle), uintptr(unsafe.Pointer(&csbi)))
	return &Writer{out: out, handle: uintptr(handle), orgAttr: uintptr(csbi.attributes)}
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
	switch code {
	case 'm':
		w.applySgr(arg)
	default:
		buf = append(buf, string(code)...)
		fmt.Fprint(w.out, string(buf))
	}
}

// Apply SGR (Select Graphic Rendition)
func (w *Writer) applySgr(arg string) {
	if arg == "" {
		procSetConsoleTextAttribute.Call(w.handle, w.orgAttr)
		return
	}
}
