package greeting

type telugu struct{}

func (e telugu) Greet() string {
	return "నమస్కరం"
}

func babu() Greeter {
	return telugu{}
}
