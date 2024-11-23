package main

import "fmt"

func main() {
	// Declare and input a float variable
	var f float32
	fmt.Printf("Enter a floating point number: ")
	fmt.Scan(&f)

	// Convert float to int and print it
	x := int(f)
	fmt.Println(x)
}
