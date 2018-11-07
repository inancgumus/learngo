// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
)

// ---------------------------------------------------------
// EXERCISE: Sphere Area
//
//  1. Get the radius from the command-line
//  2. Convert it to a float64
//  3. Calculate the surface area of a sphere
//
// SPHERE SURFACE AREA FORMULA
//  area = 4πr²
//  https://en.wikipedia.org/wiki/Sphere#Surface_area
//
// RESTRICTION
//  Use `math.Pow` to calculate the area
//  Read its documentation to see how it works.
//  https://golang.org/pkg/math/#Pow
//
// EXPECTED OUTPUT
//  go run main.go 10
//    1256.64
//
//  go run main.go 54.2
//    36915.47
// ---------------------------------------------------------

func main() {
	var radius, area float64

	// ADD YOUR CODE HERE
	// ...

	// DO NOT TOUCH THIS
	fmt.Printf("radius: %g -> area: %.2f\n", radius, area)
}
