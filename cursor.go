// +build !windows

package ansi

import (
	"fmt"
)

// Move cursor horizontally to x.
func CursorHorizontalAbsolute(x int) {
	fmt.Print("\x1b[%dG", x)
}
