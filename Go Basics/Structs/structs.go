package main

import (
	"fmt"
)

type Doctor struct {
	number int
	actorName string
	companions []string
}

func main() {
	aDoctor := Doctor{
		number: 3,
		actorName: "Jon Pertwee",
		companions: []string {
			"Liz Shaw",
			"Jo Grant",
			"Jane Smith",
		},
	}
	fmt.Println(aDoctor.companions[1])


	bDoctor := struct{name string}{name: "John Pertwee"} // anonymous struct rarely used
	bnotherDoctor := bDoctor // refers to independent structs; &bDoctor serves as the pointer
	bnotherDoctor.name = "Tom Baker"
	fmt.Println(bDoctor)
	fmt.Println(bnotherDoctor)
}