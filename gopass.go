package gopass

var g *Generator

func init() {
	g = New()
}

func Generate(length int) []byte {
	return g.Generate(length)
}

func GenerateString(length int) string {
	return g.GenerateString(length)
}

func With(opts ...Option) {
	g.With(opts...)
}
