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
	// lesson: error handling + short decl. assignment

	if len(os.Args) != 3 {
		fmt.Println("Usage: calc <number1> <number2>")
		return
	}

	a, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Please provide a valid number")
		return
	}

	b, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Println("Please provide a valid number")
		return
	}

	// err1 := ...
	// err2 := ...
	// if err1 != nil || err2 != nil {
	// 	fmt.Println("Please provide a valid number")
	// 	return
	// }

	fmt.Printf("%v + %v = %v\n", a, b, a+b)
}
