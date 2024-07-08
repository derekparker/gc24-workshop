package main

import (
	"fmt"
	"os"
)

// sayhi greets the person with the given name.
func sayhi(name string) {
	fmt.Println("Hello " + name)
}

func main() {
	// Let's say a friendly greeting
	name := os.Args[1]
	sayhi(name)

	// Let's try some arithmetic
	x := len(os.Args[1])
	if x*1 != 2 {
		panic("Wrong answer!")
	}
	fmt.Println("Goodbye")
}
