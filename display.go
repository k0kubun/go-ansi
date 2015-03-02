// +build !windows

package ansi

func EraseInLine(mode int) {
	fmt.Fprint("\x1b[%dK", x)
}
