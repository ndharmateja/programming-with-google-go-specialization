package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func read_string(scanner bufio.Scanner, prompt string) string {
	fmt.Printf(prompt)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	// Read name and address
	scanner := bufio.NewScanner(os.Stdin)
	name := read_string(*scanner, "Enter name: ")
	address := read_string(*scanner, "Enter address: ")

	// Put data in map
	data := map[string]string{"name": name, "address": address}

	// Convert to JSON and print it
	barr, _ := json.Marshal(data)
	fmt.Println("JSON data:", string(barr))
}
