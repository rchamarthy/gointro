package main

import (
	"fmt"

	"github.com/rchamarthy/gointro/greeting"
)

func main() {
	g := greeting.New("chinese")
	fmt.Println(g.Greet())
}
