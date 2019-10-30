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

// ---------------------------------------------------------
// EXERCISE: Sphere Volume
//
//  1. Get the radius from the command-line
//  2. Convert it to a float64
//  3. Calculate the volume of a sphere
//
// SPHERE VOLUME FORMULA
//  https://en.wikipedia.org/wiki/Sphere#Enclosed_volume
//
// RESTRICTION
//  Use `math.Pow` to calculate the volume
//
// EXPECTED OUTPUT
//  go run main.go 10
//    4188.79
//
//  go run main.go .5
//    0.52
// ---------------------------------------------------------

func main() {
	var radius, vol float64

	// ADD YOUR CODE HERE
	// ...

	// DO NOT TOUCH THIS
	fmt.Printf("radius: %g -> volume: %.2f\n", radius, vol)
}
