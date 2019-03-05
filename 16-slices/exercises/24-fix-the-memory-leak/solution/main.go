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

	//
	// SOLUTION #1:
	//

	// clone the last 10 elements of the returned temperatures
	// into a new slice
	need := make([]int, 10)
	copy(need, temps[len(temps)-10:])
	// make the temp slice lose reference to its backing array
	// so that it can be cleaned from the memory
	temps = need

	//
	// SOLUTION #2:
	//

	// The code below does the same thing like the code above but in one line.
	// temps = append([]int(nil), temps[len(temps)-10:]...)

	api.Report()
	fmt.Fprintln(ioutil.Discard, temps[0])
}
