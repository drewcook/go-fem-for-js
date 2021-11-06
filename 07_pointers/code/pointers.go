package main

import (
	"fmt"
	"strings"
)

func main() {
	var name string // pointer variable value, normal type
	var namePointer *string // pointer type - only accepts address since *type

	fmt.Println("Name:", name)
	fmt.Println("Name *:", namePointer)

	name = "Beyo7nce" // value for variable
	namePointer = &name // address for variable - use when have the value but need address
	var nameValue = *namePointer // value at address - use when have address but need value

	fmt.Println("Name:", name)
	fmt.Println("Name *:", namePointer)
	fmt.Println("Name Value:", nameValue)

	me := "Elvis"
	fmt.Println(me)
	changeName(&me) // use address and modify the underlying value directly
	fmt.Println(me)

	var c = Coordinates{X: 10, Y: 20}
	fmt.Println(c)
	cAddress := c
	cAddress.X = 200
	fmt.Println(cAddress)

}

// Takes in an address to a string to capitalize it and update the underlying value
func changeName(n *string) {
	*n = strings.ToUpper(*n)
}

// Passing around just an address is way cheaper to do throughout an application, as opposed to entire data sets of many bytes

// Represents latitude and longitude
type Coordinates struct {
	X, Y float64
}
