// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package api

var temps = []int{5, 10, 3, 25, 45, 80, 90}

// Read returns a slice of elements from the temps slice.
func Read(start, stop int) []int {
	// ----------------------------------------
	// RESTRICTIONS — ONLY ADD YOUR CODE IN THIS AREA

	portion := temps[start:stop]

	// RESTRICTIONS — ONLY ADD YOUR CODE IN THIS AREA
	// ----------------------------------------

	return portion
}

// All returns the temps slice
func All() []int {
	return temps
}
