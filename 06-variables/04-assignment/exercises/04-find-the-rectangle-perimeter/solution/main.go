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
	var (
		perimeter     int
		width, height = 5, 6
	)

	// first calculates: (width + height)
	// then            :  multiplies it with 2

	// just like in math

	perimeter = 2 * (width + height)

	fmt.Println(perimeter)
}
