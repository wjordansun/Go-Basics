package main

import (
	"fmt"
)

func main() {
	a := [3]int{1, 2, 3}
	b := a
	fmt.Println(a, b)
	a[1] = 42
	fmt.Println(a, b)
	// Prints 	[1 2 3] [1 2 3]
	// 			[1 42 3] [1 2 3]


	c := []int{1, 2, 3} // slice means that b refers to a
	d := c
	fmt.Println(c, d)
	c[1] = 42
	fmt.Println(c, d)
	// Prints	[1 2 3] [1 2 3]
	//			[1 42 3] [1 42 3]


	m := map[string]string {"foo": "bar", "baz": "buz"}
	n := m // maps are also reference type
	fmt.Println(m, n)
	m["foo"] = "qux"
	fmt.Println(m, n)
}