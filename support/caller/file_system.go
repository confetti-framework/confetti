package caller

import (
	"path/filepath"
	"runtime"
)

func Path() string {
	return filename()
}

func CurrentDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return filepath.Dir(filename)
}

func filename() string {
	_, filename, _, _ := runtime.Caller(2)
	return filename
}
