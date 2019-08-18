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
//  WARNING! This is a very difficult exercise. You need to
//  do some research on your own to solve it. Please don't
//  get discouraged if you can't solve it yet.
//
//  Imagine that you receive millions of temperature data
//  points but you only need the last 10 data points
//  (temperatures).
//
//  Problem: There is a memory leak in your program.
//  Please find the leak and fix it.
//
//  Memory leak means: Your program uses computer memory
//  unnecessarily. See this: https://en.wikipedia.org/wiki/Memory_leak
//
//
// CURRENT OUTPUT
//
//   > Memory Usage: 113 KB
//   > Memory Usage: 65651 KB
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
//   In the current program, because of the memory leak,
//   the difference is huge: about ~60 MB. Run the program and,
//   see it yourself.
//
//   Your goal is reducing the memory usage.
//
//   See the code in api/api.go to see how it allocates
//   a huge memory.
//
// ---------------------------------------------------------

func main() {
	// reports the initial memory usage
	api.Report()

	// reads 65 MB of temperature data into memory!
	temps := api.Read()

	// -----------------------------------------------------
	// ✪ ONLY ADD YOUR CODE INSIDE THIS BOX ✪
	//

	//
	// ✪ ONLY ADD YOUR CODE INSIDE THIS BOX ✪
	// -----------------------------------------------------

	// dont touch this code.
	api.Report()

	// don't worry about this code yet.
	fmt.Fprintln(ioutil.Discard, temps[0])
}
