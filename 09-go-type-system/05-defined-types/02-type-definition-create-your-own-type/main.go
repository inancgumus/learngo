// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
)

// type definitions usually declared at the package level
//
// EXERCISE: Move the declaration into main()'s scope
//
type (
	gram  float64 // float64 is the underlying type of gram
	ounce float64 // float64 is the underlying type of ounce
)

// above code is the same as the following:
// type gram  float64
// type ounce float64

func main() {
	// type definitions are also allowed in blocks

	// underlying types are the same
	var g gram = 1000 // gram  -> float64
	var o ounce       // ounce -> float64

	// when the underlying types are the same
	// then they can be converted between:
	// gram, ounce or float64

	// afterwards, you'll see also that,
	// the important thing is the structure
	// of the type. not just its name.
	//
	// float64 has the real structure, representation,
	// and size.
	//
	// so, it gives the structure to the newly defined type.

	// TYPE ERROR: ounce and grams are different types
	// o = g * 0.035274

	// BUT: They're convertable to each other
	o = ounce(g) * 0.035274

	fmt.Printf("%g grams is %.2f ounce\n", g, o)
}
