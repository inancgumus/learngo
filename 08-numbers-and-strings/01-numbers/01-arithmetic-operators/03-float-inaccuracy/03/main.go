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
	// When you use a float value with an integer
	// in a calculation,
	// the result always becomes a float.

	ratio := 3.0 / 2

	// OR:
	// ratio = 3 / 2.0

	// OR - if 3 is inside an int variable:
	// n := 3
	// ratio = float64(n) / 2

	fmt.Printf("%f", ratio)
}
