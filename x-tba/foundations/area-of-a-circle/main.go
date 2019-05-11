// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		radius = 10.
		area   float64
	)

	area = math.Pi * radius * radius

	fmt.Printf("radius: %g -> area: %.2f\n",
		radius, area)

	// ALTERNATIVE:
	// math.Pow calculates the power of a float number
	// area = math.Pi * math.Pow(radius, 2)
}
