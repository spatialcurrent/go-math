[![CircleCI](https://circleci.com/gh/spatialcurrent/go-math/tree/master.svg?style=svg)](https://circleci.com/gh/spatialcurrent/go-math/tree/master) [![Go Report Card](https://goreportcard.com/badge/spatialcurrent/go-math)](https://goreportcard.com/report/spatialcurrent/go-math)  [![GoDoc](https://godoc.org/github.com/spatialcurrent/go-math?status.svg)](https://godoc.org/github.com/spatialcurrent/go-math) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=math)](https://github.com/spatialcurrent/go-math/blob/master/LICENSE)

# go-math

# Description

**go-math** is a simple library that provides math functions that support varied types, but makes smart decisions to maximize precision and maintain input types.  For example, you can sum a slice of ints and floats.

**go-math** can be used in scenarios when you do not know the type of the input data a-priori.  For example, a sensor might report altitude as an integer or as a float64.

# Usage

**Go**

You can import **go-math** as a library with:

```go
import (
  "github.com/spatialcurrent/go-math/pkg/math"
)
```

See [math](https://godoc.org/github.com/spatialcurrent/go-math/pkg/math) in GoDoc for information on how to use Go API.

# Examples

See the many [examples](https://godoc.org/github.com/spatialcurrent/go-math/pkg/math) in GoDoc or the tests.

# Testing

Run test using `make test` or (`bash scripts/test.sh`), which runs unit tests, `go vet`, `go vet with shadow`, [errcheck](https://github.com/kisielk/errcheck), [ineffassign](https://github.com/gordonklaus/ineffassign), [staticcheck](https://staticcheck.io/), and [misspell](https://github.com/client9/misspell).

# Contributing

[Spatial Current, Inc.](https://spatialcurrent.io) is currently accepting pull requests for this repository.  We'd love to have your contributions!  Please see [Contributing.md](https://github.com/spatialcurrent/go-math/blob/master/CONTRIBUTING.md) for how to get started.

# License

This work is distributed under the **MIT License**.  See **LICENSE** file.
