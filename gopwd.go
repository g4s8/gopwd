package gopwd

import (
	"os"
	"path/filepath"
)

// Abs - return absolute path
// to current working directory.
// E.g. `/tmp/test` if you run app here.
func Abs() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}

// Name - return current working
// directory name, e.g. `test` if you
// run app in `/tmp/test`.
func Name() (string, error) {
	abs, err := Abs()
	if err != nil {
		return "", err
	}
	name := filepath.Base(abs)
	return name, nil
}
