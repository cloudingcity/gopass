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
			optFunc: Numbers,
		},
		"with lower case": {
			pattern: `^[a-z]+$`,
			optFunc: LowerCase,
		},
		"with upper case": {
			pattern: `^[A-Z]+$`,
			optFunc: UpperCase,
		},
		"with letters": {
			pattern: `^[a-zA-Z]+$`,
			optFunc: Letters,
		},
		"with symbols": {
			pattern: `^[[~!@#$%^&*()_+{}|:<>?,./\]]+$`,
			optFunc: Symbols,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			g := &Generator{}
			g.With(tt.optFunc())

			assert.Regexp(t, tt.pattern, g.chars.String())
		})
	}

	t.Run("with string", func(t *testing.T) {
		g := &Generator{}
		g.With(String("abc123!@#"))

		assert.Regexp(t, `^[123abc!@#]+$`, g.chars.String())
	})

	t.Run("with bytes", func(t *testing.T) {
		g := &Generator{}
		g.With(Bytes([]byte("abc123!@#")))

		assert.Regexp(t, `^[123abc!@#]+$`, g.chars.String())
	})
}
