package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name   string `required max:"100"`
	Origin string
}
type Bird struct {
	//Embedding
	Animal
	SpeedKPH float32
	CanFly   bool
}

func main() {
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b)
	fmt.Println(b.Name)

	c := Bird{
		Animal:   Animal{Name: "Cassoary", Origin: "Philipenes"},
		SpeedKPH: 32,
		CanFly:   false,
	}

	fmt.Println(c)

	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
	// tag
}
