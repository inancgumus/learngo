// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package api

import (
	"fmt"
	"math/rand"
	"runtime"
)

// DO NOT TOUCH THIS FILE BUT YOU CAN READ IT

// Read returns a huge slice (allocates ~65 MB of memory)
func Read() []int {
	// 2 << 22 means 2^(22 + 1)
	// See this: https://en.wikipedia.org/wiki/Arithmetic_shift

	// Perm function returns a slice with random integers in it.
	// Here it returns a slice with random integers that contains
	// 8,388,608 elements. One int value is 8 bytes.
	// So: 8,388,608 * 8 = ~65MB
	return rand.Perm(2 << 22)
}

// Report cleans the memory and prints the current memory usage
// Don't worry about this code. You don't need to understand it.
//
// However, if you're curious, read on.
//
// The following code runs the garbage collector to clean
// up the allocated resources, and then it reads the current
// memory statistics into the m variable.
func Report() {
	var m runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m)
	fmt.Printf(" > Memory Usage: %v KB\n", m.Alloc/1024)
}
