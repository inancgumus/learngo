// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"io/ioutil"

	"github.com/inancgumus/learngo/16-slices/exercises/24-fix-the-memory-leak/solution/api"
)

func main() {
	// reports the initial memory usage
	api.Report()

	// reads 65 MB of temperature data into the memory!
	temps := api.Read()

	// ------------------------------------------------------
	// SOLUTION #1:
	// ------------------------------------------------------

	//
	// Copy the last 10 elements of the returned temperatures
	// to a new slice.
	//
	// This will create a new backing array.
	//
	need := make([]int, 10)
	copy(need, temps[len(temps)-10:])

	//
	// Make the temp slice lose reference to its backing array
	// so that its backing array can be cleaned from the memory.
	//
	temps = need

	// ------------------------------------------------------
	// SOLUTION #2:
	// ------------------------------------------------------

	// Similar to the 1st solution. It does the same thing.
	// But this code is more concise. Use this one.

	// temps = append([]int(nil), temps[len(temps)-10:]...)

	// ------------------------------------------------------
	// don't worry about this code yet.
	api.Report()
	fmt.Fprintln(ioutil.Discard, temps[0])
}
