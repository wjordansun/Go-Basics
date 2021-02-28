package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil { // recover() is only useful in deferred functions
			log.Println("Error:", err)
			panic(err) // causes main function to panic and not print "end"
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking")
}