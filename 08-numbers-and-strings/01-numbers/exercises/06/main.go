// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE
//  Simplify the code (refactor)
//
// RESTRICTION
//  Use only the incdec and assignment operations
//
// EXPECTED OUTPUT
//  3
// ---------------------------------------------------------

func main() {
	width, height := 10, 2

	width = width + 1
	width = width + height
	width = width - 1
	width = width - height
	width = width * 20
	width = width / 25
	width = width % 5

	fmt.Println(width)
}
