package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3}
	b := a // slices automatically means b refers/points to a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))


	v := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	w := v[:]   //slice of all elements
	x := v[3:]  //slice from 4th element to end
	y := v[:6]  //slice first 6 elements
	z := v[3:6] //slice the 4th, 5th, 6th elements
	fmt.Println(v) // [1 2 3 4 5 6 7 8 9 10]
	fmt.Println(w) // [1 2 3 4 5 6 7 8 9 10]
	fmt.Println(x) // [4 5 6 7 8 9 10]
	fmt.Println(y) // [1 2 3 4 5 6]
	fmt.Println(z) // [4 5 6]

	
				  //len  cap
	m := make([]int, 3, 100)
	fmt.Println(m)
	fmt.Printf("Length: %v\n", len(m))
	fmt.Printf("Capacity: %v\n", cap(m))
	m = append(m, 1) // add to slice
	fmt.Printf("Length: %v\n", len(m))
	fmt.Printf("Capacity: %v\n", cap(m))


	p := []int{1, 2, 3, 4, 5}
	q := p[1:]
	fmt.Println(q) // slices off first element [2, 3, 4, 5]
	r := p[:len(p)-1]
	fmt.Println(r) // slices off last element [1, 2, 3, 4]
	s := append(p[:2], p[3:]...)
	fmt.Println(s) // slices off the middle element [1, 2, 4, 5]

}