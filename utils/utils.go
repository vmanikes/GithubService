// Package utils contains basic utility functions
package utils

import (
	"runtime"
	"strings"
)

// GetFuncName is used internally to determine the function name when adding it
// to the context logging.
func GetFuncName() string {
	pc := make([]uintptr, 2)
	n := runtime.Callers(1, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	frame, _ = frames.Next()

	flds := strings.Split(frame.Function, ".")
	if len(flds) >= 2 {
		return flds[len(flds)-2] + "." + flds[len(flds)-1]
	}

	return frame.Function
}
