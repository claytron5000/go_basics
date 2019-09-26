package main

import "fmt"

func main() {
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"California":   39250017,
		"Texas":        27899999,
		"Florida":      20600000,
		"New York":     19745000,
		"Pennsylvania": 19222222,
		"Illinois":     12800000,
		"Ohio":         11600000,
	}
	fmt.Println(statePopulations)
	m := map[[3]int]string{}
	fmt.Println(m)

	fmt.Println(statePopulations["Ohio"])

	statePopulations["Georgia"] = 10310371
	fmt.Println(statePopulations)

	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations)
	fmt.Println(statePopulations["Georgia"])

	// Fat finger check if presents
	_, ok := statePopulations["Oho"]
	fmt.Println(ok)

	fmt.Println(len(statePopulations))

	// By reference

	sp := statePopulations
	delete(sp, "Ohio")
	fmt.Println(sp)
	fmt.Println(statePopulations)

}
