package main

import (
	"fmt"
)

func main() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))
}

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct {} 

func (cw ConsoleWriter) Write(data []byte) (int, error) { // implementation of interface
	n, err := fmt.Println(string(data))
	return n, err
}