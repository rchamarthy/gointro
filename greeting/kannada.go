package greeting

type kannada struct{}

func (t kannada) Greet() string {
	return "ನಮಸ್ಕಾರ"
}

func maga() Greeter {
	return kannada{}
}
