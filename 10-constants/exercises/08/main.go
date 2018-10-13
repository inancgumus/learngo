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
//  1. Initialize multiple constants using iota.
//  2. Please follow the instructions inside the code.
//
// EXPECTED OUTPUT
//  1 2 3
// ---------------------------------------------------------

func main() {
	// HINT: This is a valid constant declaration
	//       Blank-Identifier can be used in place of a name
	const _ = iota
	//    ^- this is just a name

	// Now, use iota and initialize the following constants
	// "automatically" to 1, 2, and 3 respectively.
	const (
		Jan = iota
		Feb
		Mar
	)

	// This should print: 1 2 3
	// Not: 0 1 2
	fmt.Println(Jan, Feb, Mar)
}
