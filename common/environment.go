package common

import (
	"bytes"
	"io/ioutil"
	"io"
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

// StandardInput is an abstraction on os.Stdin
func StandardInput() io.Reader {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return os.Stdin
	}

	if stat.Size() > 0 {
		contents, _ := ioutil.ReadFile(os.Stdin.Name())
		return bytes.NewReader(contents)
	}

	return os.Stdin
}
