package gopass

import (
	"bytes"
	"crypto/rand"
	"math/big"
)

var DefaultOptions = []Option{Numbers(), Letters()}

type Generator struct {
	chars bytes.Buffer
}

func New(opts ...Option) Gopass {
	g := &Generator{}
	g.With(opts...)
	return g
}

func (g *Generator) Generate(length int) []byte {
	if length <= 0 {
		return nil
	}

	buf := newBuffer()
	defer releaseBuffer(buf)

	g.generate(buf, length)

	return buf.Bytes()
}

func (g *Generator) GenerateString(length int) string {
	if length <= 0 {
		return ""
	}

	buf := newBuffer()
	defer releaseBuffer(buf)

	g.generate(buf, length)

	return buf.String()
}

func (g *Generator) With(opts ...Option) {
	g.chars.Reset()

	if len(opts) == 0 {
		opts = DefaultOptions
	}
	for _, opt := range opts {
		opt(g)
	}
}

func (g *Generator) generate(buf *bytes.Buffer, length int) {
	max := big.NewInt(int64(g.chars.Len()))
	for i := 0; i < length; i++ {
		n, _ := rand.Int(rand.Reader, max)
		buf.WriteByte(g.chars.Bytes()[n.Int64()])
	}
}
