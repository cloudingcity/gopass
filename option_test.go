package gopass

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptions(t *testing.T) {
	tests := map[string]struct {
		pattern string
		optFunc func() Option
	}{
		"with numbers": {
			pattern: `^[0-9]+$`,
			optFunc: WithNumbers,
		},
		"with lower case": {
			pattern: `^[a-z]+$`,
			optFunc: WithLowerCase,
		},
		"with upper case": {
			pattern: `^[A-Z]+$`,
			optFunc: WithUpperCase,
		},
		"with letters": {
			pattern: `^[a-zA-Z]+$`,
			optFunc: WithLetters,
		},
		"with symbols": {
			pattern: `^[[~!@#$%^&*()_+{}|:<>?,./\]]+$`,
			optFunc: WithSymbols,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			g := &Generator{}
			g.options(tt.optFunc())

			assert.Regexp(t, tt.pattern, g.chars.String())
		})
	}

	t.Run("with chars", func(t *testing.T) {
		g := &Generator{}
		g.options(WithChars("abc123!@#"))

		assert.Regexp(t, `^[123abc!@#]+$`, g.chars.String())
	})
}
