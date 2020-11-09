package gopass

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	g := New()

	t.Run("default", func(t *testing.T) {
		got := g.Generate(10)
		assert.Len(t, got, 10)
		assert.Regexp(t, "^[0-9a-zA-z]+$", string(got))
	})

	t.Run("invalid length", func(t *testing.T) {
		got := g.Generate(-10)
		assert.Nil(t, got)
	})
}

func TestGenerateString(t *testing.T) {
	g := New()

	t.Run("default", func(t *testing.T) {
		got := g.GenerateString(10)
		assert.Len(t, got, 10)
		assert.Regexp(t, "^[0-9a-zA-z]+$", got)
	})

	t.Run("invalid length", func(t *testing.T) {
		got := g.GenerateString(-10)
		assert.Empty(t, got)
	})
}

func BenchmarkGenerate(b *testing.B) {
	g := New()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			g.Generate(100)
		}
	})
}

func BenchmarkGenerateString(b *testing.B) {
	g := New()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			g.GenerateString(100)
		}
	})
}
