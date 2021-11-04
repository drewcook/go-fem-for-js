package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "John"
	names[1] = "Sam"
	names[2] = "Greg"
	var scores [5]float32 = [5]float32{9, 1.2, 4.5, 8, 10.02}
	nums := [...]int{2, 3, 4, 5, 6, 7, 8, 9, 0}

	fmt.Println(names)
	fmt.Println(scores)
	fmt.Println(nums)


	// iterate
	for _, s := range scores {
		fmt.Println("Score:", s)
	}
}
