// +build !windows

package ansi

// Move cursor horizontally to x.
func CursorHorizontalAbsolute(x int) {
	fmt.Fprint("\x1b[%dG", x)
}
