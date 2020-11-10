package gopass

const (
	numbers   = "0123456789"
	lowerCase = "abcdefghijklmnopqrstuvwxyz"
	upperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letters   = lowerCase + upperCase
	symbols   = "~!@#$%^&*()_+{}|:<>?,./"
)

type Option func(*Generator)

func Numbers() Option {
	return func(g *Generator) {
		g.chars.WriteString(numbers)
	}
}

func UpperCase() Option {
	return func(g *Generator) {
		g.chars.WriteString(upperCase)
	}
}

func LowerCase() Option {
	return func(g *Generator) {
		g.chars.WriteString(lowerCase)
	}
}

func Letters() Option {
	return func(g *Generator) {
		g.chars.WriteString(letters)
	}
}

func Symbols() Option {
	return func(g *Generator) {
		g.chars.WriteString(symbols)
	}
}

func String(s string) Option {
	return func(g *Generator) {
		g.chars.WriteString(s)
	}
}

func Bytes(b []byte) Option {
	return func(g *Generator) {
		g.chars.Write(b)
	}
}
