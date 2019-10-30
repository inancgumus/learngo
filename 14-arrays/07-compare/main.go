// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// STORY:
// You want to compare two bookcases,
// whether they're equal or not.

func main() {
	// When comparing two arrays, their types should be identical

	// Comment out one of the following pair of variables and observe the results.

	var (
		// equal (types + elements are identical)::
		blue = [3]int{6, 9, 3}
		red  = [3]int{6, 9, 3}

		// equal (types + elements are identical):
		// blue = [...]int{6, 9, 3}
		// red  = [3]int{6, 9, 3}

		// not equal (element ordering are different):
		// blue = [3]int{6, 9, 3}
		// red  = [3]int{3, 9, 6}

		// not equal (the last elements are not equal):
		// blue = [3]int{6, 9}
		// red  = [3]int{6, 9, 3}

		// not comparable (type mismatch: length):
		// blue = [3]int{6, 9, 3}
		// red  = [5]int{6, 9, 3}

		// not comparable (type mismatch: element type):
		// blue = [3]int64{6, 9, 3}
		// red  = [3]int{6, 9, 3}
	)
	fmt.Printf("blue bookcase : %v\n", blue)
	fmt.Printf("red bookcase  : %v\n", red)

	fmt.Println("Are they equal?", blue == red)
}
