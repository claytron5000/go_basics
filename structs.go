package main

import "fmt"

type Doctor struct {
	number     int
	actorName  string
	companions []string
	episodes   []string
}

func main() {
	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Jane Smith",
		},
	}

	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.companions[1])

	// Anonymous struct
	anotherDoctor := struct{ name string }{name: "Howard Hughs"}
	againAnotherDoctor := anotherDoctor
	againAnotherDoctor.name = "Tom Baker"
	fmt.Println(anotherDoctor)
	fmt.Println(againAnotherDoctor)
}
