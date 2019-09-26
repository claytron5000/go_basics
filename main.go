package main

import (
	"fmt"
	"strconv"
)

var I int = 44

func main() {
	i := 43
	fmt.Printf("%v, %T\n", i, i)

	var j string
	j = strconv.Itoa(i)

	fmt.Printf("%v, %T", j, j)
}
