package gopwd

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

// Abs - return absolute path
// to current working directory.
// E.g. `/tmp/test` if you run app here.
func Abs() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error on read $PWD")
	}
	return dir, nil
}

// Name - return current working
// directory name, e.g. `test` if you
// run app in `/tmp/test`.
func Name() (string, error) {
	dir, err := Abs()
	if err != nil {
		return "", err
	}
	var ss []string
	if runtime.GOOS == "windows" {
		ss = strings.Split(dir, "\\")
	} else {
		ss = strings.Split(dir, "/")
	}

	currentDirName := ss[len(ss)-1]
	return currentDirName, nil
}
