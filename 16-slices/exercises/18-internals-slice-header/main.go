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
//  Note that, your memory usage numbers may vary. These are on my
//  own system. However, the size of the arrays and slices should be
//  the same on your own system as well (if you're on 64-bit machine).
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
//    - Accepts a [size]int array, so you can pass it your array.
//    - It automatically prints the memory usage.
//
//    passSlice function:
//    - Accepts an int slice, so you can pass it one of your slices.
//    - It automatically prints the memory usage.
//
// ---------------------------------------------------------

const size = 1e7

func main() {
	// stops the gc: prevents cleaning up the memory
	debug.SetGCPercent(-1)

	// run the program to see what this prints
	report("initial memory usage")

	// 1. allocate an array with 10 million int elements
	//    this array's size is equal to ~80MB
	//    hint: use the `size` constant
	//
	// 2. print the memory usage

	// 3. copy the array to a new array (just assign)
	// 4. print the memory usage

	// 5. pass the array to the passArray function

	// 6. convert the one of the arrays to a slice (by slicing)
	// 7. slice only the first 1000 elements of the array
	// 8. slice only the elements of the array between 1000 and 10000
	// 9. print the memory usage

	// 10. pass the one of the slices to the passSlice function

	// 11. print the sizes of the arrays and slices
	//     hint: use the unsafe.Sizeof function
}

// observe that passing an array affects the memory usage dramatically
//
// passes [size]int array — about 80MB!
func passArray(items [size]int) {
	items[0] = 100
	report("inside passArray")
}

// observe that passing a slice doesn't affect the memory usage
//
// only passes 24-bytes of slice header
func passSlice(items []int) {
	items[0] = 100
	report("inside passSlice")
}

func report(msg string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("[%s]\n", msg)
	fmt.Printf("\t> Memory Usage: %v KB\n", m.Alloc/1024)
}
