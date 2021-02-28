package main

import (
	"fmt"
	"reflect" //reflection package used for tags
)

type Animal struct{
	Name 	string `required max: "100"` // backtick is used for tags
	Origin 	string
}

type Bird struct{
	Animal
	SpeedKPH float32
	CanFly bool
}

func main() {
	b := Bird{} // instantiate
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b.Name)

	bi := Bird{
		Animal: Animal{Name: "Emu", Origin: "Australia"},
		SpeedKPH: 48,
		CanFly: false,
	}
	fmt.Println(bi.Name)

	t := reflect.TypeOf(Animal{}) // gets type of function
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)

}