package common

import (
	"os"
	"path/filepath"
)

// CurrentExecutable gets the path to this executable.
func CurrentExecutable() string {
	fn := os.Args[0]
	fd := filepath.Dir(fn)
	fp, _ := filepath.Abs(fd)
	return fp + "/" + fn
}
