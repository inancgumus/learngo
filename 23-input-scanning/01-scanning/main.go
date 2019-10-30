// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Simulate an error
	// os.Stdin.Close()

	// Create a new scanner that scans from the standard-input
	in := bufio.NewScanner(os.Stdin)

	// Stores the total number of lines in the input
	var lines int

	// Scan the input line by line
	for in.Scan() {
		lines++

		// Get the current line from the scanner's buffer
		// fmt.Println("Scanned Text :", in.Text())
		// fmt.Println("Scanned Bytes:", in.Bytes())
		in.Text()
	}
	fmt.Printf("There are %d line(s)\n", lines)

	if err := in.Err(); err != nil {
		fmt.Println("ERROR:", err)
	}
}
