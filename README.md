# go-envparser

[![Go Reference](https://pkg.go.dev/badge/github.com/m0t0k1ch1/go-envparser.svg)](https://pkg.go.dev/github.com/m0t0k1ch1/go-envparser)
![test](https://github.com/m0t0k1ch1/go-envparser/workflows/test/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/m0t0k1ch1/go-envparser/badge.svg?branch=main)](https://coveralls.io/github/m0t0k1ch1/go-envparser?branch=main)

parse environment variables as the specified Go type

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

func main() {
	os.Setenv("PORT", "12345")

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
- float32
- float64
