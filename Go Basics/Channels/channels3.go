package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 50)
	wg.Add(2)
	go func(ch <-chan int) { // recieve only channel
		for i := range ch {
			if i, ok := <- ch; ok; { // ok is a boolean saying whether the channel is open
				fmt.Println(i)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)
	go func(ch chan<- int) { // send only channel
		ch <- 42
		ch <- 27
		close(ch) // closes channel to stop for-loop
		wg.Done()
	}(ch)
	wg.Wait()
}