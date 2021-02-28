package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")
	//Prints	start
	//			end
	//			middle

	// defer fmt.Println("start")
	// defer fmt.Println("middle")
	// defer fmt.Println("end")
	// Prints	end
	//			middle
	//			start
	// Based on Last In First Out (LIFO)


	//a := "start"
	//defer fmt.Println(a)  // takes in parameter at time of defer
	//a = "end"
	//Prints start

}