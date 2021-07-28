package greeting

type tamil struct{}

func (t tamil) Greet() string {
	return "வணக்கம்"
}

func ponnu() Greeter {
	return tamil{}
}
