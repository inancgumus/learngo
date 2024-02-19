// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Fix
//
//  1. Uncomment the code
//
//  2. And fix the problems
//
//  3. BONUS: Simplify the code
// ---------------------------------------------------------

func main() {
	names := [...]string{"Einstein", "Shepard", "Tesla"}
	books := [5]string{"Kafka's Revenge", "Stay Golden"}

	fmt.Printf("%#v\n", names)
	fmt.Printf("%#v\n", books)
}
