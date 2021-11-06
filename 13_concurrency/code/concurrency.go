package main

import (
	"fmt"
	"sync"
	"time"
)

// Setup a WaitGroup
var wg sync.WaitGroup

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("We panicked, but everything is fine. Message received: ", r)
	}
}

func say(s string) {
	// Defer
	defer wg.Done()
	defer handlePanic()

	for i := 0; i < 10; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 700)
	}
}

func sayLonger(s string) {
	// Defer
	defer wg.Done() // releses a count off the stack of wg
	defer handlePanic()

	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 1500)
		if i == 2 {
			panic("Uh Oh!")
		}
	}
}

func main() {
	// Increment wait group counter
	wg.Add(1)
	// Launch a goroutine
	go say("Hello")

	// Let's do some longer stuff too
	wg.Add(1)
	go sayLonger("Thinking...")

	// Increment wait group counter again
	wg.Add(1)
	// Launch another goroutine
	go say("There")

	// Wait for the wait group counter to be 0 before continuing
	wg.Wait()
	fmt.Println("Running code after goroutines complete...")
}
