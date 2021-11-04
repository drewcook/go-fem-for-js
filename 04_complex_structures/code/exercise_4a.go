package main

import "fmt"

// variadic function
func average(nums ...float32) float32 {
	var sum float32
	for _, num := range nums {
		sum += num
	}
	return sum / float32(len(nums)) // len returns int, cast it to float for math
}


func main() {
	avg := average(120.32, 28.3, 487, 39, 0)
	fmt.Println(avg)
}
