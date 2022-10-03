// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

func main() {
	min := int8(127)
	max := int16(1000)

	fmt.Println(max + int16(min))

	// EXPLANATION
	//
	// `int8(max)` destroys the information of max
	// It reduces it to 127
	// Which is the maximum value of int8
	//
	// Correct conversion is int16(min)
	// Because, int16 > int8
	// When you do so, min doesn't lose information
	//
	// You will learn more about this in
	// the "Go Type System" section.
}
