# go-envparser

[![Go Report Card](https://goreportcard.com/badge/github.com/m0t0k1ch1/go-envparser)](https://goreportcard.com/report/github.com/m0t0k1ch1/go-envparser)
[![Go Reference](https://pkg.go.dev/badge/github.com/m0t0k1ch1/go-envparser.svg)](https://pkg.go.dev/github.com/m0t0k1ch1/go-envparser)
[![test](https://github.com/m0t0k1ch1/go-envparser/actions/workflows/test.yml/badge.svg)](https://github.com/m0t0k1ch1/go-envparser/actions/workflows/test.yml)
[![Coverage Status](https://coveralls.io/repos/github/m0t0k1ch1/go-envparser/badge.svg?branch=main)](https://coveralls.io/github/m0t0k1ch1/go-envparser?branch=main)

A tiny package to parse environment variables as the specified Go type.

## Installation

```sh
$ go get github.com/m0t0k1ch1/go-envparser
```

## Example

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/m0t0k1ch1/go-envparser"
)

func init() {
	os.Setenv("PORT", "12345")
}

func main() {
	var port uint16
	if err := envparser.Parse("PORT", &port); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("PORT: %d", port)
}
```

## Supported types

- string
- int
- int8
- int16
- int32
- int64
- uint
- uint8
- uint16
- uint32
- uint64
- bool
  - `false` for `""` or `"0"`, `true` otherwise
