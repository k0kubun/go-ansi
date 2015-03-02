// +build !windows

package ansi

import (
	"fmt"
)

// Move cursor horizontally to x.
func CursorHorizontalAbsolute(x int) {
	fmt.Printf("\x1b[%dG", x)
}
