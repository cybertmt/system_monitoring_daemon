package utils

import (
	"fmt"
	"runtime"
)

var ErrNotImplemented = fmt.Errorf("statistics not implemented for: %s", runtime.GOOS)
