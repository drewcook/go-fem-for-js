package main

import "fmt"

// part 1 - arrays
func partOne() {
	fmt.Println("===== Part One =====")
	var scores [5]float64 = [5]float64{38.92, 45.84, 29.84, 74.29, 90.02}
	fmt.Println(scores)
	res := getAverage(scores)
	fmt.Println(res)
}
func getAverage(nums [5]float64) float64 {
	sum := 0.0
	for _, n := range nums {
		sum += n
	}
	return sum / float64(len(nums))
}

// part 2 - maps
func partTwo() {
	fmt.Println("===== Part Two =====")
	pets := map[string]string{
		"fido": "dog",
		"fluffy": "cat",
	}
	pets["rosky"] = "dog"
	pets["sherlock"] = "fish"
	fmt.Println(pets)
	r1 := hasPet(pets, "rosky")
	r2 := hasPet(pets, "john")
	fmt.Println(r1, r2)
}
func hasPet(pets map[string]string, name string) bool {
	if pettype, ok := pets[name]; ok {
		fmt.Printf("Yes, you have a %s named %s.\n", pettype, name)
		return true
	} else {
		fmt.Printf("Sorry, you do not have a pet named %s.\n", name)
	}
	return false
}

// part 3 - slices
var groceries = []string{"butter", "bananas", "milk", "tofu"}
func partThree() {
	fmt.Println("===== Part 3 =====")
	slice := groceries[1:3]
	moreGroceries := append(slice, "nuts", "tortillas", "hummas")
	fmt.Println(groceries, moreGroceries)
	list := addGrocery("jam", "kale", "salsa")
	fmt.Println(list)
}
func addGrocery(items ...string) []string {
	newList := groceries
	for _, item := range items {
		newList = append(groceries, item)
	}
	return newList
}

func main() {
	partOne()
	partTwo()
	partThree()
}

