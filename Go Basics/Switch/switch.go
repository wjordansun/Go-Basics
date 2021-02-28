package main

import (
	"fmt"
)

func main() {		// tag
	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("one, five, or ten")
	case 2, 4, 6:
		fmt.Println("two, four, or six")
	default:
		fmt.Println("another number")
	}


	i := 10
	switch { // tagless switches allow cases to overlap
	case i <= 10:
		fmt.Println("less than or equal to ten")
		fallthrough // logic less. executes statements in next case regardless
	case i <= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than twenty")
	}


	var a interface{} = 1
	switch a.(type) {
	case int:
		fmt.Println("i is an int")
		break // break out early
		fmt.Println("This will print too")
	case float64:
		fmt.Println("i is  a float64")
	case string:
		fmt.Println("i is string")
	default:
		fmt.Println("i is another type")
	}
}