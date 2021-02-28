package main

import (
	"fmt"
)

func main() {
	statePopulations := map[string]int { // basically a dictionary; key and value
		"California":	39250017,
		"Texas": 	  	27862596,
		"Florida":	  	20612439,
		"New York":	  	19745289,
		"Pennsylvania":	12802503,
		"Illinois":		12801539,
		"Ohio":			11614373,

	}

	statePopulations["Georgia"] = 10310371
	fmt.Println(statePopulations)
	
	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations)

	fmt.Println(statePopulations["Ohio"])

	fmt.Println(len(statePopulations))
	
	pop, ok := statePopulations["Ohio"] // ok is a booleann that checks whether "Ohio" is in
	fmt.Println(pop, ok)				// the map

	sp := statePopulations // refers to statePopulations
	fmt.Println(sp)

	// Another way to declare maps
	// statePopulations := make(map[string]int)
	// statePopulations = map[string]int { // basically a dictionary; key and value
	// 	"California":	39250017,
	// 	"Texas": 	  	27862596,
	// 	"Florida":	  	20612439,
	// 	"New York":	  	19745289,
	// 	"Pennsylvania":	12802503,
	// 	"Illinois":		12801539,
	// 	"Ohio":			11614373,

	// }
}