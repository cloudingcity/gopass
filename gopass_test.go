package gopass

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGopass(t *testing.T) {
	t.Run("generate", func(t *testing.T) {
		got := Generate(10)
		assert.Len(t, got, 10)
		assert.Regexp(t, "^[0-9a-zA-Z]+$", string(got))
	})

	t.Run("generate string", func(t *testing.T) {
		got := GenerateString(10)
		assert.Len(t, got, 10)
		assert.Regexp(t, "^[0-9a-zA-Z]+$", got)
	})

	t.Run("generate string with customize characters", func(t *testing.T) {
		With(String("ABC"))
		got := GenerateString(10)
		assert.Len(t, got, 10)
		assert.Regexp(t, "^[A-C]+$", got)
	})
}
