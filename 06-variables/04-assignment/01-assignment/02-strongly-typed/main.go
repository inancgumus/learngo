// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// Go is a strongly typed programming language.

// Even a float and an integer are different types.
// Or: int32 and int are different types.

// EXERCISE: Try uncommenting the lines
//           And observe the errors

func main() {
	var speed int
	// speed = "100"

	var running bool
	// running = 1

	var force float64
	// speed = force

	// Go automatically converts the typeless
	//   integer literal to float64 automatically
	force = 1

	// keep the compiler happy
	_, _, _ = speed, running, force
}
