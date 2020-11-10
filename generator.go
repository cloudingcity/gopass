package gopass

import (
	"bytes"
	"math/rand"
	"time"
)

var DefaultOptions = []Option{WithNumbers(), WithLetters()}

type Generator struct {
	chars bytes.Buffer
}

func New(opts ...Option) *Generator {
	g := &Generator{}
	g.options(opts...)
	return g
}

func (g *Generator) Generate(length int) []byte {
	if length <= 0 {
		return nil
	}

	buf := NewBuffer()
	defer ReleaseBuffer(buf)

	g.generate(buf, length)

	return buf.Bytes()
}

func (g *Generator) GenerateString(length int) string {
	if length <= 0 {
		return ""
	}

	buf := NewBuffer()
	defer ReleaseBuffer(buf)

	g.generate(buf, length)

	return buf.String()
}

func (g *Generator) options(opts ...Option) {
	if len(opts) == 0 {
		opts = DefaultOptions
	}
	for _, opt := range opts {
		opt(g)
	}
}

func (g *Generator) generate(buf *bytes.Buffer, length int) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < length; i++ {
		buf.WriteByte(g.chars.Bytes()[rand.Intn(g.chars.Len())])
	}
}
