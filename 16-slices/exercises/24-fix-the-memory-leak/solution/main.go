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
	"io/ioutil"

	"github.com/inancgumus/learngo/16-slices/exercises/24-fix-the-memory-leak/solution/api"
)

func main() {
	// reports the initial memory usage
	api.Report()

	// returns a slice with 10 million elements.
	// it allocates 65 MB of memory space.
	millions := api.Read()

	// ------------------------------------------------------

	// SOLUTION #1:
	// Copy the last 10 elements of the returned slice
	// to a new slice. This will create a new backing array
	// only with 10 elements.
	last10 := make([]int, 10)
	copy(last10, millions[len(millions)-10:])

	// Make the millions slice lose reference to its backing array
	// so that its backing array can be cleaned up from memory.
	millions = last10

	// SOLUTION #2:
	// Similar to the 1st solution. It does the same thing.
	// But this code is more concise. Use this one.

	// millions = append([]int(nil), millions[len(millions)-10:]...)

	fmt.Printf("\nLast 10 elements: %d\n\n", last10)

	// ------------------------------------------------------

	api.Report()

	// don't worry about this code yet.
	fmt.Fprintln(ioutil.Discard, millions[0])
}
