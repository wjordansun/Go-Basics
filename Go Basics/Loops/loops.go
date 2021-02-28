package main

import (
	"fmt"
)

func main() {
	for i, j := 0, 0; i < 5; i, j = i + 1, j + 2 {
		fmt.Println(i, j)
	}


	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}


	i := 0 //scoped to main function
	for i < 5 {
		fmt.Println(i)
		i++
	}


	j := 0
	for {
		fmt.Println(j)
		j++
		if j == 5 {
			break
		}
	}


	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

Loop:

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++{
			fmt.Println(i * j)
			if i * j >= 3 {
				break Loop
			}
		}
	}


	s := []int{1, 2, 3}
	for k, v := range s { // k is index, v is value
		fmt.Println(k, v)
	}


	statePopulations := map[string]int {
		"California":	39250017,
		"Texas": 	  	27862596,
		"Florida":	  	20612439,
		"New York":	  	19745289,
		"Pennsylvania":	12802503,
		"Illinois":		12801539,
		"Ohio":			11614373,

	}
	for k, v := range statePopulations {
		fmt.Println(k, v)
	}


	t := "hello Go!"
	for k, v := range t {
		fmt.Println(k, string(v))
	}

}