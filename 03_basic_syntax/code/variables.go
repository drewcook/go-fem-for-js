package main

import (
	"fmt"
	"reflect"
)

var dog string
func main() {
	var name = "Jimbo"
	dog = "Lucy"
	// similar to 'let' keyword in JS
	var lastname string
	var someNumber int
	cat := "Garfield"
	fmt.Println(name, reflect.TypeOf(name))
	fmt.Println(someNumber, reflect.TypeOf(someNumber))
	fmt.Println(lastname, reflect.TypeOf(lastname))
	fmt.Println(dog, cat)
}
