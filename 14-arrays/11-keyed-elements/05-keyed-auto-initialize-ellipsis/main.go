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
	// below ellipsis (...) calculates the length of the
	// array automatically
	rates := [...]float64{
		// index 0 empty
		// index 1 empty
		// index 2 empty
		// index 3 empty
		// index 4 empty
		5: 1.5, // index: 5
	}

	fmt.Println(rates)

	// above array literal equals to this:
	//
	// rates := [6]float64{
	// 	0.,
	// 	0.,
	// 	0.,
	// 	0.,
	// 	0.,
	// 	1.5,
	// }
}
