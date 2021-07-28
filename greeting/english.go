package greeting

type english struct{}

func (e english) Greet() string {
	return "hello"
}

func englishMan() Greeter {
	return english{}
}
