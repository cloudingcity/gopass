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

	b := NewBuffer()
	defer ReleaseBuffer(b)

	g.generate(b, length)

	return b.Bytes()
}

func (g *Generator) GenerateString(length int) string {
	if length <= 0 {
		return ""
	}

	b := NewBuffer()
	defer ReleaseBuffer(b)

	g.generate(b, length)

	return b.String()
}

func (g *Generator) options(opts ...Option) {
	if len(opts) == 0 {
		opts = DefaultOptions
	}
	for _, opt := range opts {
		opt(g)
	}
}

func (g *Generator) generate(b *bytes.Buffer, length int) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < length; i++ {
		b.WriteByte(g.chars.Bytes()[rand.Intn(g.chars.Len())])
	}
}
