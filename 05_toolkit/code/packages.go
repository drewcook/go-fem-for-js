package main

import (
	"fmt"
	"go-fem-for-js/05_toolkit/code/utils"
)

func calcData() int {
	totalValue := utils.Add(28, 38, 92, 30, 11)
	return totalValue
}

func main() {
	fmt.Println("Packages!")
	total := calcData()
	fmt.Println(total)
}
