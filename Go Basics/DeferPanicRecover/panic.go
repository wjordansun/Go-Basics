package main

import (
	"fmt"
)

func main() {
	// a, b := 1, 0
	// ans := a / b
	// fmt.Println(ans)
	// Panic runtime error

	fmt.Println("start")
	panic("something bad happened")
	fmt.Println("end")
	
}