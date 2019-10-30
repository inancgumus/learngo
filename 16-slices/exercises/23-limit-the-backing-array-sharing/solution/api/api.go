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
	//
	// The third index prevents the `main()` from
	// overwriting the original temps slice's
	// backing array. It limits the capacity of the
	// returned slice. See the full slice expressions
	// lecture for more details.
	//
	portion := temps[start:stop:stop]

	return portion
}

// All returns the temps slice
func All() []int {
	return temps
}
