package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{} // wait group syncs go routines so they happen at the same time
						//	 channels synchorinize the data passed betweenn them
func main() {
	ch := make(chan int) // channel intialization
	wg.Add(2)
	go func(ch <-chan int) { // recieve only channel
		i := <- ch
		fmt.Println(i)
		wg.Done()
	}(ch)
	go func(ch chan<- int) { // send only channel
		i := 42
		ch <- i
		wg.Done()
	}(ch) // These goroutines can be embedded in a for loop
	wg.Wait()
}