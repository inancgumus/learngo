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
	i := 42 + 3.14 // OK: 42 and 3.14 are typeless

	// above line equals to this:
	// i := float64(42 + 3.14)

	fmt.Printf("i is %T\n", i)
}
