// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	// `safe`'s value is `false`
	var safe bool

	// `safe`'s value is now `true`

	// `speed` is declared and initialized to `50`

	// redeclaration only works when
	//
	// at least one of the variables
	// should be a new variable

	safe, speed := true, 50

	fmt.Println(safe, speed)

	// EXERCISE
	//
	// Declare the speed variable before
	// the short declaration "again"
	//
	// Observe what happens
}
