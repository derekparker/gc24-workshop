package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// This program accepts input from
// stdin and determins if the input is an even or odd number
// and prints the results.
func main() {
	fmt.Println("Pid:", os.Getpid())
	buf := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Enter a number: ")
		str, err := buf.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		str = strings.TrimSuffix(str, "\n")
		n, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
		if n%2 == 0 {
			fmt.Println("The number is even!")
		} else {
			fmt.Println("The number is odd!")
		}
	}
}
