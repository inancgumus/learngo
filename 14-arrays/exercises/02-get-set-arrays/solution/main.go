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
	"strings"
)

func main() {
	var (
		names     [3]string  // The names of your three best friends
		distances [5]int     // The distances to five different locations
		data      [5]byte    // A data buffer with five bytes of capacity
		ratios    [1]float64 // Currency exchange ratios only for a single currency
		alives    [4]bool    // Up/Down status of four different web servers
		zero      [0]byte    // A byte array that doesn't occupy memory space
	)

	names[0] = "Einstein"
	names[1] = "Tesla"
	names[2] = "Shepard"

	distances[0] = 50
	distances[1] = 40
	distances[2] = 75
	distances[3] = 30
	distances[4] = 125

	data[0] = 'H'
	data[1] = 'E'
	data[2] = 'L'
	data[3] = 'L'
	data[4] = 'O'

	ratios[0] = 3.14145

	alives[0] = true
	alives[1] = false
	alives[2] = true
	alives[3] = false

	// zero[0] = "BOMB!"
	_ = zero

	// =========================================================================

	separator := "\n" + strings.Repeat("=", 20) + "\n"

	fmt.Print("names", separator)
	for i := 0; i < len(names); i++ {
		fmt.Printf("names[%d]: %q\n", i, names[i])
	}

	fmt.Print("\ndistances", separator)
	for i := 0; i < len(distances); i++ {
		fmt.Printf("distances[%d]: %d\n", i, distances[i])
	}

	fmt.Print("\ndata", separator)
	for i := 0; i < len(data); i++ {
		// try the %c verb
		fmt.Printf("data[%d]: %d\n", i, data[i])
	}

	fmt.Print("\nratios", separator)
	for i := 0; i < len(ratios); i++ {
		fmt.Printf("ratios[%d]: %.2f\n", i, ratios[i])
	}

	fmt.Print("\nalives", separator)
	for i := 0; i < len(alives); i++ {
		fmt.Printf("alives[%d]: %t\n", i, alives[i])
	}

	// no loop for zero elements
	fmt.Print("\nzero", separator)
	for i := 0; i < len(zero); i++ {
		fmt.Printf("zero[%d]: %d\n", i, zero[i])
	}

	// =========================================================================

	// you know how this works :) don't be freaked out!
	fmt.Printf(`

%s
FOR RANGES
%[1]s

`, strings.Repeat("~", 30))

	fmt.Print("names", separator)
	for i, v := range names {
		fmt.Printf("names[%d]: %q\n", i, v)
	}

	fmt.Print("\ndistances", separator)
	for i, v := range distances {
		fmt.Printf("distances[%d]: %d\n", i, v)
	}

	fmt.Print("\ndata", separator)
	for i, v := range data {
		// try the %c verb
		fmt.Printf("data[%d]: %d\n", i, v)
	}

	fmt.Print("\nratios", separator)
	for i, v := range ratios {
		fmt.Printf("ratios[%d]: %.2f\n", i, v)
	}

	fmt.Print("\nalives", separator)
	for i, v := range alives {
		fmt.Printf("alives[%d]: %t\n", i, v)
	}

	// no loop for zero elements
	fmt.Print("\nzero", separator)
	for i, v := range zero {
		fmt.Printf("zero[%d]: %d\n", i, v)
	}
}
