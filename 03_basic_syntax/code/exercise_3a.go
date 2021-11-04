package main

import "fmt"

func main() {
	sentence := "I like to play guitar."

	for idx, char := range sentence {
		if idx % 2 != 0 {
			fmt.Println(string(char))
		}
	}
}
