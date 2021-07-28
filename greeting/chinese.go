package greeting

type chinese struct{}

func (c chinese) Greet() string {
	return "你好"
}

func chinaMan() Greeter {
	return chinese{}
}
