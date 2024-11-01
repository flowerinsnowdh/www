package util

import (
	"fmt"
	"os"
)

func ErrPrintln(a ...any) {
	_, _ = fmt.Fprintln(os.Stderr, a)
}
