package main

import (
	"fmt"
)

func main() {
	// Boolean
	n := 1 == 1
	m := 1 == 2

	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", m, m)

	// Signed int
	i := 42
	fmt.Printf("%v, %T\n", i, i)
	// Unsigned int
	var p uint16 = 42
	fmt.Printf("%v, %T\n", p, p)

	// Int operations
	a := 10
	b := 3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	// Type conversion cross int types
	var g int = 10
	var h int8 = 3
	fmt.Println(g + int(h))

	// Bit operators
	fmt.Println("Bit operators")
	j := 10             // 1010
	k := 3              // 0011
	fmt.Println(j & k)  // 0010
	fmt.Println(j | k)  // 1011
	fmt.Println(j ^ k)  // 1001
	fmt.Println(j &^ k) // 0100

	// bit shifting
	fmt.Println("Bit shifting")
	w := 8              // 2^3
	fmt.Println(w << 3) // 2^3 * 2^3 =  2^6
	fmt.Println(w >> 3) // 2^3 / 2^3 = 2^0

	// Floating point
	var c float64 = 3.14
	c = 13.7e72
	c = 2.14E14
	fmt.Printf("%v, %T\n", c, c)

	fmt.Println("Complet types")
	// Complex type
	var y complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", y, y)
	fmt.Printf("%v, %T\n", real(y), real(y))
	fmt.Printf("%v, %T\n", imag(y), imag(y))

	o := 1 + 2i
	l := 2 + 5.2i

	fmt.Println(o + l)
	fmt.Println(o - l)
	fmt.Println(o * l)
	fmt.Println(o / l)

	// String types
	fmt.Println("String types")

	s := "this is a string"
	fmt.Printf("%v, %T\n", string(s[2]), s[2])

	s2 := "this is also a string"
	fmt.Printf("%v, %T\n", s+s2, s+s2)

	e := []byte(s)
	fmt.Printf("%v, %T\n", e, e)

	//Rune type
	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r)

}
