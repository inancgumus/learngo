// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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
	}
	fmt.Printf("There %d lines.\n", lines)
}
