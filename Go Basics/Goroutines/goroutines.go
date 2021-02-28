package main

import (
	"fmt"
	"time"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	 go sayHello() // abstraction over a thread
	 time.Sleep(100 * time.Millisecond) // allows statement to be printed


	 var msg = "Hello"
	 wg.Add(1)
	 go func(msg string) {
	 	fmt.Println(msg)
	 	wg.Done()
	 }(msg)
	 msg = "Goodbye"
	 wg.Wait() // allows go routine to run without real world clock
}

func sayHello() {
	fmt.Println("Hello")
}