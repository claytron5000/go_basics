package main

import (
	"fmt"
)

const a int16 = 27

const (
	j = iota
	k
	l
)

const (
	m = iota
)

const (
	errorSpecialist = iota // or _ = iota if this is truely unreachable
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {
	// overridable
	const a int = 15
	const b string = "foo"
	const c float32 = 3.12
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)

	const e int = 27
	var f int = 42
	fmt.Printf("%v, %T\n", e+f, e+f)

	// inferred types
	const g = 42
	var h int16 = 27
	fmt.Printf("%v, %T\n", g+h, g+h)

	// Inumerated constants
	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v\n", j)
	fmt.Printf("%v\n", k)
	fmt.Printf("%v\n", l)
	fmt.Printf("%v\n", m)

	var specialistType int
	fmt.Printf("%v\n", specialistType == catSpecialist)

	// bit shifting iota to make powers of 2
	fileSize := 4000000000.
	fmt.Printf("%.2fGB\n", fileSize/GB)

	// bit shifting oring and masking for roles
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is Headquarters? %v\n", isHeadquarters&roles == isHeadquarters)
}
