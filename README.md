# Logur integration for [Logr](https://github.com/go-logr/logr) interface

[![CircleCI](https://circleci.com/gh/logur/integration-logr.svg?style=svg)](https://circleci.com/gh/logur/integration-logr)
[![Coverage](https://gocover.io/_badge/logur.dev/integration/logr)](https://gocover.io/logur.dev/integration/logr)
[![Go Report Card](https://goreportcard.com/badge/logur.dev/integration/logr?style=flat-square)](https://goreportcard.com/report/logur.dev/integration/logr)
[![GolangCI](https://golangci.com/badges/github.com/logur/integration-logr.svg)](https://golangci.com/r/github.com/logur/integration-logr)
[![Go Version](https://img.shields.io/badge/go%20version-%3E=1.11-61CFDD.svg?style=flat-square)](https://github.com/logur/integration-logr)
[![GoDoc](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/logur.dev/integration/logr)


## Installation

```bash
go get logur.dev/integration/logr
```


## Usage

```go
package main

import (
	"github.com/goph/logur"
	logrintegration "logur.dev/integration/logr"
)

func main() {
	logger := logrintegration.New(logur.NewNoopLogger())
}
```


## Development

When all coding and testing is done, please run the test suite:

``` bash
$ make check
```


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
