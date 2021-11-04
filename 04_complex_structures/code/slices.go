package main

import "fmt"

func main() {

	var myArray [5]int
	var mySlice []int = make([]int, 3, 10)

	myArray[0] = 1
	mySlice[0] = 1
	mySlice[2] = 1

	// fmt.Println(myArray)
	// fmt.Println(mySlice)
	// fmt.Println(len(mySlice))
	// fmt.Println(cap(mySlice))

	// fmt.Println(len(myArray))
	// fmt.Println(cap(myArray))

	// ***************************

	// var myArray [5]int
	// var mySlice []int = make([]int, 5)

	// fmt.Println(myArray)
	// fmt.Println(mySlice)

	// ***************************

	// var myArray [5]int
	// // var mySlice []int = make([]int, 5)
	// var mySlice []int = make([]int, 5, 10)
	// // var mySlice = make([]int, 5, 10)

	// myArray[0] = 1
	// mySlice[0] = 1

	// fmt.Println(myArray)
	// fmt.Println(mySlice)
	// fmt.Println(len(mySlice))
	// fmt.Println(cap(mySlice))

	// ***************************

	fruitArray := [5]string{"banana", "pear", "apple", "kumquat", "peach"}
	slicedFruit := fruitArray[1:3] // gets index 1 up to but not including 3, so index 1 and 2
	fmt.Println(slicedFruit) // [pear apple]
	fmt.Println(len(slicedFruit), cap(slicedFruit)) // 2 4
	// adding to array next
	withMoreFruit := append(slicedFruit, "grape", "stawberry", "lemon", "lime", "orange", "plantain", "kiwi", "pomelo")
	withEvenMoreFruit := append(withMoreFruit, "blackberry") // adding more than original capacity can hold will double it
	fmt.Println(slicedFruit, withMoreFruit)
	fmt.Println(len(withMoreFruit), cap(withMoreFruit))
	fmt.Println(len(withEvenMoreFruit), cap(withEvenMoreFruit))

	// ***************************

	// SEE SLIDE

	// ***************************

	// slice1 := []int{1, 2, 3}
	// slice2 := append(slice1, 4, 5)

	// fmt.Println(slice1, slice2)
	// fmt.Println(len(slice1), cap(slice1))
	// fmt.Println(len(slice2), cap(slice2))

	// ***************************

	// originalSlice := []int{1, 2, 3}
	// destination := make([]int, len(originalSlice))

	// fmt.Println("Before Copy:", originalSlice, destination)

	// mysteryValue := copy(destination, originalSlice)

	// // fmt.Println("After Copy:", originalSlice, destination, mysteryValue)
}
