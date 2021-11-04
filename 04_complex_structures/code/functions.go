package main

import "fmt"

func printAge(age, year int) (ageOfBob int, ageOfSally int) {
	// do something with age/year
	// declare return vars
	ageOfBob = 29
	ageOfSally = 13
	return
}

func printAges(ages ...int) { // spread, however many args are passed in, print them
	for _, age := range ages {
		fmt.Println(age)
	}
}

func main() {
	r1, r2 := printAge(8, 2021)
	fmt.Println(r1, r2)
	printAges(100, 200, 300, 400)
}
