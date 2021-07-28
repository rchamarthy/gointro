package greeting

type hindi struct{}

func (t hindi) Greet() string {
	return "नमस्ते"
}

func gori() Greeter {
	return hindi{}
}
