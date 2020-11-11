# Gopass

[![Test](https://github.com/cloudingcity/gopass/workflows/Test/badge.svg)](https://github.com/cloudingcity/gopass/actions?query=workflow%3ATest)
[![Lint](https://github.com/cloudingcity/gopass/workflows/Lint/badge.svg)](https://github.com/cloudingcity/gopass/actions?query=workflow%3ALint)
[![codecov](https://codecov.io/gh/cloudingcity/gopass/branch/main/graph/badge.svg)](https://codecov.io/gh/cloudingcity/gopass)
[![Go Report Card](https://goreportcard.com/badge/github.com/cloudingcity/gopass)](https://goreportcard.com/report/github.com/cloudingcity/gopass)

:key: Gopass is a cryptographically secure random password generator Go library.

## QuickStart

```go
package main

import "github.com/cloudingcity/gopass"

func main() {
	gopass.GenerateString(20) // vJd3NRrCImrSQf3M1h3A
}
```

## Examples

Generate string contains letters and numbers with the given length.

> Letters and numbers is the default option.

```go
gopass.GenerateString(20) // vJd3NRrCImrSQf3M1h3A
```

Generate with options.

```go
gopass.With(gopass.Numbers(), gopass.UpperCase(), gopass.Symbols())
gopass.GenerateString(20) // ^34J)GV.AE,@KA?+~MPG
```

Generate []byte
```go
gopass.Generate(20) // [90 55 81 120 104 70 114 105 101 54]
```

Customize characters.

```go
gopass.With(gopass.String("123"), gopass.Bytes([]byte("ABC")))
gopass.GenerateString(10) // 31AABCBB22
```

Create your own gopass instance.
```go
pass := gopass.New(gopass.Numbers(), gopass.UpperCase())
pass.GenerateString(20) // AMQPZJ5OUGW4WVI8GB3C
```

## Benchmarks

```console
$ GOMAXPROCS=4 go test -bench=Generate -benchmem -benchtime=3s
BenchmarkGenerate-4               146948             25280 ns/op            4808 B/op        301 allocs/op
BenchmarkGenerateString-4         167050             25397 ns/op            4920 B/op        302 allocs/op
```
