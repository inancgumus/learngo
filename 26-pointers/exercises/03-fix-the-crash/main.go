// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Fix the crash
//
// EXPECTED OUTPUT
//
//   brand: apple
// ---------------------------------------------------------

package main

import "fmt"

type computer struct {
	brand *string
}

func main() {
	var c *computer
	change(c, "apple")
	fmt.Printf("brand: %s\n", c.brand)
}

func change(c *computer, brand string) {
	(*c.brand) = brand
}
