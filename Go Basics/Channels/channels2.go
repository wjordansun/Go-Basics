package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 50)// buffer allows data to be stored in the channel
						// useful when frequencies of sending and recieving are different
	wg.Add(2)
	go func(ch <-chan int) { // recieve only channel
		i := <-ch
		fmt.Println(i)
		i = <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)
	go func(ch chan<- int) { // send only channel
		ch <- 42
		ch <- 27
		wg.Done()
	}(ch)
	wg.Wait()
}