package gopass

const (
	numbers   = "0123456789"
	lowerCase = "abcdefghijklmnopqrstuvwxyz"
	upperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letters   = lowerCase + upperCase
	symbols   = "~!@#$%^&*()_+{}|:<>?,./"
)

type Option func(*Generator)

func WithNumbers() Option {
	return func(g *Generator) {
		g.chars.WriteString(numbers)
	}
}

func WithUpperCase() Option {
	return func(g *Generator) {
		g.chars.WriteString(upperCase)
	}
}

func WithLowerCase() Option {
	return func(g *Generator) {
		g.chars.WriteString(lowerCase)
	}
}

func WithLetters() Option {
	return func(g *Generator) {
		g.chars.WriteString(letters)
	}
}

func WithSymbols() Option {
	return func(g *Generator) {
		g.chars.WriteString(symbols)
	}
}

func WithChars(s string) Option {
	return func(g *Generator) {
		g.chars.WriteString(s)
	}
}
