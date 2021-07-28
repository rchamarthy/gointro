package greeting

type sanskrit struct{}

func (s sanskrit) Greet() string {
	return "नमस्ते"
}

func purusha() Greeter {
	return sanskrit{}
}
