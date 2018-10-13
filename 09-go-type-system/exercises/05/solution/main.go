// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"os"
	"strconv"
)

// WHY?
// Because, you can add methods to your types in the future
// And, they're type-safe and more readable now.

func main() {
	type (
		Feet   float64
		Meters float64
	)

	var (
		feet   Feet
		meters Meters
	)

	arg := os.Args[1]

	// feet is a Feet value now
	val, _ := strconv.ParseFloat(arg, 64)
	feet = Feet(val)

	// meters is a Meters value now
	meters = Meters(feet * 0.3048)

	fmt.Printf("%g feet is %g meters.\n", feet, meters)
}
