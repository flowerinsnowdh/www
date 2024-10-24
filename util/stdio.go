package util

import (
    "fmt"
    "os"
)

func StderrPrintln(a... any) {
    _, _ = fmt.Fprintln(os.Stderr, a)
}
