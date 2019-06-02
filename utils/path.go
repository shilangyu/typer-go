package utils

import (
	"path"
	"runtime"
)

// Root returns the root directory of the project
func Root() string {
	_, filePath, _, _ := runtime.Caller(0)
	return path.Join(filePath, "..", "..")
}
