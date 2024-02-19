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

func main() {
	var radius, vol float64

	radius, _ = strconv.ParseFloat(os.Args[1], 64)

	vol = (4 * math.Pi * math.Pow(radius, 3)) / 3

	fmt.Printf("radius: %g -> volume: %.2f\n", radius, vol)
}
