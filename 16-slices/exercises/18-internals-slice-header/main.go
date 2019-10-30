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
	"runtime"
	"runtime/debug"
)

// ---------------------------------------------------------
// EXERCISE: Observe the memory allocations
//
//  In this exercise, your goal is to observe the memory allocation
//  differences between arrays and slices.
//
//  You will create, assign arrays and slices then you will print
//  the memory usage of your program on each step.
//
//  Please follow the instructions inside the code.
//
//
// EXPECTED OUTPUT
//
//  Note that, your memory usage numbers may vary. However, the size of the
//  arrays and slices should be the same on your own system as well
//  (if you're on a 64-bit machine).
//
//
//  [initial memory usage]
//          > Memory Usage: 104 KB
//  [after declaring an array]
//          > Memory Usage: 78235 KB
//  [after copying the array]
//          > Memory Usage: 156365 KB
//  [inside passArray]
//          > Memory Usage: 234495 KB
//  [after slicings]
//          > Memory Usage: 234497 KB
//  [inside passSlice]
//          > Memory Usage: 234497 KB
//
//  Array's size : 80000000 bytes.
//  Array2's size: 80000000 bytes.
//  Slice1's size: 24 bytes.
//  Slice2's size: 24 bytes.
//  Slice3's size: 24 bytes.
//
//
// HINTS
//
//  I've declared a few functions to help you.
//
//    report function:
//    - Prints the memory usage.
//    - Just call it with a message that matches to the expected output.
//
//    passArray function:
//    - It automatically prints the memory usage.
//    - Accepts a [size]int array, so you can pass your array to it.
//
//    passSlice function:
//    - It automatically prints the memory usage.
//    - Accepts an int slice, so you can pass it one of your slices.
//
// ---------------------------------------------------------

const size = 1e7

func main() {
	// don't worry about this code.
	// it stops the garbage collector: prevents cleaning up the memory.
	// see the link if you're curious:
	// https://en.wikipedia.org/wiki/Garbage_collection_(computer_science)
	debug.SetGCPercent(-1)

	// run the program to see the initial memory usage.
	report("initial memory usage")

	// 1. allocate an array with 10 million int elements
	//    the array's size will be equal to ~80MB
	//    hint: use the `size` constant above.

	// 2. print the memory usage (use the report func).

	// 3. copy the array to a new array.

	// 4. print the memory usage

	// 5. pass the array to the passArray function

	// 6. convert one of the arrays to a slice

	// 7. slice only the first 1000 elements of the array

	// 8. slice only the elements of the array between 1000 and 10000

	// 9. print the memory usage (report func)

	// 10. pass the one of the slices to the passSlice function

	// 11. print the sizes of the arrays and slices
	//     hint: use the unsafe.Sizeof function
	//     see more here: https://golang.org/pkg/unsafe/#Sizeof
}

// passes [size]int array — about 80MB!
//
// observe that passing an array to a function (or assigning it to a variable)
// affects the memory usage dramatically
func passArray(items [size]int) {
	items[0] = 100
	report("inside passArray")
}

// only passes 24-bytes of slice header
//
// observe that passing a slice doesn't affect the memory usage
func passSlice(items []int) {
	items[0] = 100
	report("inside passSlice")
}

// reports the current memory usage
// don't worry about this code
func report(msg string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("[%s]\n", msg)
	fmt.Printf("\t> Memory Usage: %v KB\n", m.Alloc/1024)
}
