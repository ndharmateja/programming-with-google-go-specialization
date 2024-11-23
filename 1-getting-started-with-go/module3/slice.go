package main

import (
	"fmt"
	"strconv"
)

func main() {
	slice := make([]int, 0, 3)
	for i := 0; i < 5; i++ {
		// Take input
		var r string
		fmt.Printf("Enter an integer (X to quit): ")
		fmt.Scan(&r)

		// If user enters 'x' we exit
		if r == "x" || r == "X" {
			break
		}

		// Insert it into the slice
		x, _ := strconv.Atoi(r)
		slice = append(slice, x)

		// Sort the slice
		for i := len(slice) - 1; i > 0 && slice[i] < slice[i-1]; i-- {
			temp := slice[i]
			slice[i] = slice[i-1]
			slice[i-1] = temp
		}

		// Print the slice
		fmt.Println(slice)
	}
}
