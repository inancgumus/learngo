// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

// I didn't talk about this in a lecture
// As a side-note, I wanted to put it here
// Please review it.

var declareMeAgain = 10

func nested() { // block scope starts

	// declares the same variable
	// they both can exist together
	// this one only belongs to this scope
	// package's variable is still intact
	var declareMeAgain = 5
	fmt.Println("inside nope:", declareMeAgain)

} // block scope ends

func main() { // block scope starts

	fmt.Println("inside main:", declareMeAgain)

	nested()

	// package-level declareMeAgain isn't effected
	// from the change inside the nested func
	fmt.Println("inside main:", declareMeAgain)

} // block scope ends
