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

	"github.com/inancgumus/learngo/16-slices/exercises/24-fix-the-memory-leak/api"
)

// ---------------------------------------------------------
// EXERCISE: Fix the memory leak
//
//  You receive millions of temperature data points from
//  an API, but you only need the last 10 data points.
//
//  Currently, there is a memory leak in your program.
//  Find the leak and fix it.
//
//
// EXPECTED OUTPUT
//
//   > Memory Usage: 116 KB
//   > Memory Usage: 118 KB
//
//
// EXPECTED OUTPUT EXPLANATION
//
//   Your output will be different. Your goal is to reduce the
//   difference between two measurements of the memory usage.
//
//   For the expected output above:
//
//     118 KB - 116 KB = Only 2 KB so that's OK.
//
//   However, in the current program, because of the memory leak,
//   the difference is huge: about ~60 MB. Run the program and,
//   see yourself.
//
// ---------------------------------------------------------

func main() {
	// reports the initial memory usage
	api.Report()

	// reads 65 MB of temperature data into the memory!
	temps := api.Read()

	// -----------------------------------------------------
	// ✪ ONLY ADD YOUR CODE INSIDE THIS BOX ✪
	//

	//
	// ✪ ONLY ADD YOUR CODE INSIDE THIS BOX ✪
	// -----------------------------------------------------

	// fix the problem so that the memory usage stays low
	// dont touch this code
	api.Report()
	fmt.Fprintln(ioutil.Discard, temps[0])
}
