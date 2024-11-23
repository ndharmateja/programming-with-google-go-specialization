package main

import (
	"fmt"
	"strings"
)

func main() {
	// Take string input
	var input string
	fmt.Printf("Enter a string: ")
	fmt.Scan(&input)

	// Convert string to lower case
	input = strings.ToLower(input)

	if input[0] == 'i' && input[len(input)-1] == 'n' && strings.Contains(input, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
