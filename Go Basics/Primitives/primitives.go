package main

import (
	"fmt"
)

func main() {
	var test bool = true //default initialization of 0, or false
	fmt.Printf("%v, %T\n", test, test)

	n := 1 == 1
	m := 1 == 2
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T", m, m)


//	var i uint16 = 42 //unsigned integer means positive; uint8 is from 0 - 255, where int8 
					  // is from -128 - 127
	fmt.Printf("%v, %T\n", n ,n)

	a := 10 // 1010
	b := 3  // 0011

	fmt.Println(a + b) 
	fmt.Println(a - b) 
	fmt.Println(a * b) 
	fmt.Println(a / b)
	fmt.Println(a % b)

	fmt.Println(a & b) // 0010 = 2
	fmt.Println(a | b) // 1011 = 11
	fmt.Println(a ^ b) // 1001 = 9
	fmt.Println(a &^ b)// 0100 = 8

	// bit shifting
	c := 8 // 2^3
	fmt.Println(c << 3) // 2^3 * 2^3 = 2^6
	fmt.Println(c >> 3) // 2^3 / 2^3 = 2^0


	var d complex64 = 1 + 2i // imaginary number
	fmt.Printf("%v, %T\n", real(d), real(d))
	fmt.Printf("%v, %T\n", imag(d), imag(d))


	s := "this is a string"
	fmt.Printf("%v, %T\n", string(s[2]), s[2])

	sb := []byte(s)
	fmt.Printf("%v, %T\n", sb, sb)

	r := 'a' // rune, int32
	fmt.Printf("%v, %T\n", r, r)


}