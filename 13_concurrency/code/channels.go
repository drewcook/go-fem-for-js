package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func doMath(c chan int, val int) {
	defer wg.Done()
	c <- val * 5 // adds value into channel concurrently
}

func main() {
	// second arg is buffer to tell channel how many times we plan on using it
	resultChan := make(chan int, 10)

	// Fire off goroutines
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doMath(resultChan, i)
	}

	// Wait for all goroutines to complete before closing the channel
	wg.Wait()

	// Close the channel
	close(resultChan)

	// Print out results from our channel
	for item := range resultChan {
		fmt.Println(item)
	}
}
