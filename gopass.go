package gopass

func init() {
	g = New()
}

type Gopass interface {
	Generate(length int) []byte
	GenerateString(length int) string
	With(opts ...Option)
}

var g Gopass

func Generate(length int) []byte {
	return g.Generate(length)
}

func GenerateString(length int) string {
	return g.GenerateString(length)
}

func With(opts ...Option) {
	g.With(opts...)
}
