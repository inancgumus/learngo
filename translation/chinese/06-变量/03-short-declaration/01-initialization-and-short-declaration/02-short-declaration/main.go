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
	// OPTION #1 (option #2 is better)
	// var safe bool = true

	// OPTION #2 (OK)
	// var safe = true

	// OPTION #3 - SHORT DECLARATION (BEST)
	//
	// You don't even need to type the `var` keyword
	//
	// Short declaration equals to:
	//   var safe bool = true
	//   var safe = true
	//
	// Go gets (infers) the type from the initializer value
	//
	// true's default type is bool
	// so, the type of the safe variable becomes a bool
	safe := true

	fmt.Println(safe)
}
