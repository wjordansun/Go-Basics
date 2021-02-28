package main

import (
	"fmt"
)

func main() {
	var a int = 42
	var b *int = &a  // address of operator
	fmt.Println(a, *b) // 42, 42 // *b means dereferencing
	a = 27
	fmt.Println(a, *b) // 27, 27
	*b = 14
	fmt.Println(a, *b) // 14, 14

	i := [3]int{1, 2, 3}
	j := &i[0]
	k := &i[1]
	fmt.Printf("%v %p %p\n", i, j, k)


	var ms *myStruct
	ms = &myStruct{foo: 42}
	fmt.Println(ms) // Prints &{42}

	// var ms *myStruct
	// ms = new(myStruct)
	// (*ms).foo = 42 // the * is not needed b/c Go interprets it
	// fmt.Println((*ms).foo) // Prints 42



}

type myStruct struct {
	foo int
}