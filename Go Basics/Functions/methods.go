package main

import (
	"fmt"
)

func main() {
	g := greeter {
		greeting: "Hello",
		name: "Go",
	}
	g.greet()
	fmt.Println("The new name is:", g.name)
}

type greeter struct {
	greeting string
	name string
}

		//pointer allows to manipulate underlying data
func (g *greeter) greet() {
	fmt.Println(g.greeting, g.name)
	g.name = ""
}