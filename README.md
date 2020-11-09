# Gopass

[![Test](https://github.com/cloudingcity/gopass/workflows/Test/badge.svg)](https://github.com/cloudingcity/gopass/actions?query=workflow%3ATest)
[![Lint](https://github.com/cloudingcity/gopass/workflows/Lint/badge.svg)](https://github.com/cloudingcity/gopass/actions?query=workflow%3ALint)
[![codecov](https://codecov.io/gh/cloudingcity/gopass/branch/main/graph/badge.svg)](https://codecov.io/gh/cloudingcity/gopass)
[![Go Report Card](https://goreportcard.com/badge/github.com/cloudingcity/gopass)](https://goreportcard.com/report/github.com/cloudingcity/gopass)

:key: Gopass is a thread-safe random password generator Go library and zero memory allocation.

## QuickStart

```go
package main

import "github.com/cloudingcity/gopass"

func main() {
	pass := gopass.New()
	pass.GenerateString(20) // vJd3NRrCImrSQf3M1h3A
}
```

## Benchmarks

```console
$ GOMAXPROCS=4 go test -bench=Generate -benchmem -benchtime=10s
BenchmarkGenerate-4               613816             22660 ns/op               0 B/op          0 allocs/op
BenchmarkGenerateString-4         565225             22211 ns/op             112 B/op          1 allocs/op
```

## Examples

Generate string contains letters and numbers with the given length.

```go
pass := gopass.New()
pass.GenerateString(20) // vJd3NRrCImrSQf3M1h3A
```

Generate with options.

```go
pass := gopass.New(gopass.WithNumbers(), gopass.WithUpperCase())
pass.GenerateString(20) // AMQPZJ5OUGW4WVI8GB3C
```

Generate []byte
```go
pass := gopass.New()
pass.Generate(20) // [90 55 81 120 104 70 114 105 101 54]
```

Customize characters.

```go
pass := gopass.New(gopass.WithString("123"), gopass.WithBytes([]byte("ABC")))
pass.GenerateString(10) // 31AABCBB22
```
