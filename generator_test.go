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
		assert.Len(t, got, 0)
		assert.Nil(t, got)
	})
}

func TestGenerateString(t *testing.T) {
	g := New()
	got := g.GenerateString(10)
	assert.Len(t, got, 10)
	assert.Regexp(t, "^[0-9a-zA-z]+$", got)
}

func BenchmarkGenerate(b *testing.B) {
	g := New()
	for i := 0; i < b.N; i++ {
		g.Generate(30)
	}
}

func BenchmarkGenerateString(b *testing.B) {
	g := New()
	for i := 0; i < b.N; i++ {
		g.GenerateString(30)
	}
}
