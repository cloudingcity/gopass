package gopass

import (
	"bytes"
	"math/rand"
	"time"
)

var DefaultOptions = []Option{WithNumbers(), WithNumbers()}

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

	rand.Seed(time.Now().UnixNano())

	b := NewBuffer()
	defer ReleaseBuffer(b)

	for i := 0; i < length; i++ {
		b.WriteByte(g.randByte())
	}
	return b.Bytes()
}

func (g *Generator) GenerateString(length int) string {
	if length <= 0 {
		return ""
	}

	rand.Seed(time.Now().UnixNano())

	b := NewBuffer()
	defer ReleaseBuffer(b)

	for i := 0; i < length; i++ {
		b.WriteByte(g.randByte())
	}
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

func (g *Generator) randByte() byte {
	return g.chars.Bytes()[rand.Intn(g.chars.Len())]
}
