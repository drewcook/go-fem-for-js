package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

// This throws an error if the number is less than 10.
func greaterThanTen(n int) error {
	if n < 10 {
		msg := fmt.Sprintf("ERROR: %d is not greater than 10!", n)
		return errors.New(msg)
	}
	return nil
}

// This throws an error if can't find file to open
func openFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err // error type from os.Open()
	}
	defer f.Close()
	return nil
}

func handlePanic() string {
	return "HANDLING THE PANIC"
}

func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("We panicked but everything is fine.")
		fmt.Println("Recovery instructions:", r)
	}
}

func panicOnTwo() {
	defer recoverFromPanic()
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i == 2 {
			panic(handlePanic())
		}
	}
}

func doThings() {
  defer fmt.Println("First Line but do this last!") // multiple defers are LIFO
  defer fmt.Println("Do this second to last!")
  fmt.Println("Things And Stuff should happen first")
}

func main() {
	// Basic errors
	// err := greaterThanTen(higher)
	if err := greaterThanTen(10); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("All good")
	}

	if err := greaterThanTen(8); err != nil {
		fmt.Println(err) // continues to run subsequent code
		// panic(err) // breaks the program (stops it)
		// log.Fatalln(err) // great for logging these errors somewhere with timestamps. set up in our logging configs
	} else {
		fmt.Println("All good")
	}

	// Open File
	if err := openFile("/missing.txt"); err != nil {
		fmt.Println(fmt.Errorf("%v", err))
	}

	// Panicking
	doThings()
	panicOnTwo()
}
