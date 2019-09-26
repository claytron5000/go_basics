package main

import (
	"fmt"
)

func main() {

	a := []int{1, 2, 3}
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

	// slices are pointers by nature
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	aa := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	bb := aa[:]
	c := aa[3:]
	d := aa[:6]
	e := aa[3:6]
	aa[5] = 99
	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	f := make([]int, 3, 100)
	fmt.Println(f)
	fmt.Printf("Length: %v\n", len(f))
	fmt.Printf("Capacity: %v\n", cap(f))

	g := []int{}
	fmt.Println(g)
	fmt.Printf("Length: %v\n", len(g))
	fmt.Printf("Capacity: %v\n", cap(g))
	g = append(g, 1)
	fmt.Println(g)
	fmt.Printf("Length: %v\n", len(g))
	fmt.Printf("Capacity: %v\n", cap(g))
	g = append(g, 2, 3, 4, 5)
	fmt.Println(g)
	fmt.Printf("Length: %v\n", len(g))
	fmt.Printf("Capacity: %v\n", cap(g))

	// Concatenate a slice
	g = append(g, []int{6, 7, 8, 9}...)
	fmt.Println(g)
	fmt.Printf("Length: %v\n", len(g))
	fmt.Printf("Capacity: %v\n", cap(g))

	h := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// pull from begining
	i := h[1:]
	fmt.Println(i)
	// pop
	j := h[:len(h)-1]
	fmt.Println(j)
	// pluck
	fmt.Println(h)
	k := append(h[:2], h[3:]...)
	fmt.Println(k)
	// reference havock!
	fmt.Println(h)

}
