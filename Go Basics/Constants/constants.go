package main

import (
	"fmt"
)

const (
	a = iota // 0, int
	b = iota // 1, int
	c = iota // 2, int
	// iota increments
)

const (
	a2 = iota // restarts at 0
)

const (
	_ = iota // error catching
	catSpecialist // Go predicts subsequent iota declarations
	dogSpecialist
	snakeSpecialist
)

func main() {

	const myConst int = 42
	fmt.Printf("%v, %T\n", myConst, myConst) // 42, int

	const m = 42
	var n int16 = 27
	fmt.Printf("%v, %T\n", m + n, m + n) //69, int16

	fmt.Printf("%v\n", a) // 0
	fmt.Printf("%v\n", b) // 1
	fmt.Printf("%v\n", c) // 2
	fmt.Printf("%v\n", a2)// 0

	var specialistType int
	fmt.Printf("%v\n", specialistType == catSpecialist)

}