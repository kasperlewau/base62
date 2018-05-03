# base62 [![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/kasperlewau/base62) [![Build Status](https://travis-ci.org/kasperlewau/base62.svg?branch=master)](https://travis-ci.org/kasperlewau/base62) [![Go Report Card](https://goreportcard.com/badge/github.com/kasperlewau/base62)](https://goreportcard.com/report/github.com/kasperlewau/base62)
> converts ints to byteslices/strings of base62 data and back again.

## install
```sh
go get github.com/kasperlewau/base62
```

## usage 
```go
import (
	"fmt"
	
	"github.com/kasperlewau/base62"
)

func main() {
	base62.Encode(99) // [49 66]
	base62.EncodeString(99) // "1B"

	base62.Decode([]byte("1B")) // 99
	base62.DecodeString("1B") // 99
}
```

## benchmarks
```sh
goos: linux
goarch: amd64
pkg: github.com/kasperlewau/base62
BenchmarkEncoders/kasperlewau/bytes-8           30000000                34.7 ns/op             8 B/op          1 allocs/op
BenchmarkEncoders/kasperlewau/string-8          30000000                44.0 ns/op            16 B/op          1 allocs/op
BenchmarkEncoders/pilu/go-base62-8              30000000                43.5 ns/op            16 B/op          1 allocs/op
BenchmarkEncoders/marksalpeter/token-8          10000000               230 ns/op              39 B/op          2 allocs/op
BenchmarkDecoders/kasperlewau/bytes-8           50000000                32.9 ns/op             0 B/op          0 allocs/op
BenchmarkDecoders/kasperlewau/string-8          30000000                41.5 ns/op             0 B/op          0 allocs/op
BenchmarkDecoders/pilu/go-base62-8              10000000               175 ns/op              64 B/op          1 allocs/op
BenchmarkDecoders/marksalpeter/token-8          10000000               193 ns/op              64 B/op          1 allocs/op
```

## license 
MIT
