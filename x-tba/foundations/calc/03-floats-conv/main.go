// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// lesson: len(), floats encompass integers too

	if len(os.Args) != 3 {
		fmt.Println("Usage: calc <number1> <number2>")
		return
	}

	a, _ := strconv.ParseFloat(os.Args[1], 64)
	b, _ := strconv.ParseFloat(os.Args[2], 64)

	fmt.Printf("%v + %v = %v\n", a, b, a+b)
}
