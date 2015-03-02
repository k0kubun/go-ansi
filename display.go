// +build !windows

package ansi

import (
	"fmt"
)

func EraseInLine(mode int) {
	fmt.Print("\x1b[%dK", mode)
}
