[![DevOps By Rultor.com](http://www.rultor.com/b/g4s8/gopwd)](http://www.rultor.com/p/g4s8/gopwd)

[![CircleCI](https://circleci.com/gh/g4s8/gopwd.svg?style=svg)](https://circleci.com/gh/g4s8/gopwd)
[![Build Status](https://img.shields.io/travis/g4s8/gopwd.svg?style=flat-square)](https://travis-ci.org/g4s8/gopwd)

[![PDD status](http://www.0pdd.com/svg?name=g4s8/gopwd)](http://www.0pdd.com/p?name=g4s8/gopwd)
[![License](https://img.shields.io/github/license/g4s8/gopwd.svg?style=flat-square)](https://github.com/g4s8/gopwd/blob/master/LICENSE)

Golang library to get current working directory absolute path or its name.

## Usage

Download: `go get -u github.com/g4s8/gopwd`

Example:
```go
import "github.com/g4s8/gopwd"
import "fmt"

func main() {
  pwd, err := gopwd.Name()
  if err != nil {
    panic(err)
  }
  fmt.Printf("Current directory is %s\n", pwd)
}
```

## Contribution
Fork repository, clone it, make changes,
push to new branch and submit a pull request.
