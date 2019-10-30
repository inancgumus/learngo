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
	"math"
	"os"
	"strconv"
)

// This program contains a lot of comments
// Non-commented version is in /without-comments/main.go

func main() {
	// You cannot use this syntax.
	//
	// feetInMeters is declared
	//   when the declaration completes,
	//   (at the end of the whole declaration)
	//
	// And, it's not easy to read
	//
	// const feetInMeters, feetInYards float64 = 0.3048,
	//   feetInMeters / 0.9144

	const (
		feetInMeters = 0.3048
		feetInYards  = feetInMeters / 0.9144

		// cannot call runtime funcs
		// feetInYards = math.Round(feetInYards)
	)

	// Immutable: You cannot assign new values
	// feetInYards = 0.333333

	arg := os.Args[1]

	feet, _ := strconv.ParseFloat(arg, 64)

	meters := feet * feetInMeters
	yards := math.Round(feet * feetInYards)

	fmt.Printf("%g feet is %g meters.\n", feet, meters)
	fmt.Printf("%g feet is %g yards.\n", feet, yards)
}
