package utils

import "fmt"

func printNum(num int) {
	fmt.Println("Current number:", num)
}

// Add - adds together multiple numbers
func Add(nums ...int) int {
	total := 0
	for _, n := range nums {
		printNum(n)
		total += n
	}
	return total
}

