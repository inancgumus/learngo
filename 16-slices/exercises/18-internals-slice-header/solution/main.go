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
	"unsafe"
)

const size = 1e7

func main() {
	debug.SetGCPercent(-1)

	report("initial memory usage")

	var array [size]int
	report("after declaring an array")

	array2 := array
	report("after copying the array")

	passArray(array)

	slice1 := array[:]
	slice2 := array[:1e3]
	slice3 := array[1e3:1e4]
	report("after slicings")

	passSlice(slice3)

	fmt.Println()
	fmt.Printf("Array's size : %d bytes.\n", unsafe.Sizeof(array))
	fmt.Printf("Array2's size: %d bytes.\n", unsafe.Sizeof(array2))
	fmt.Printf("Slice1's size: %d bytes.\n", unsafe.Sizeof(slice1))
	fmt.Printf("Slice2's size: %d bytes.\n", unsafe.Sizeof(slice2))
	fmt.Printf("Slice3's size: %d bytes.\n", unsafe.Sizeof(slice3))
}

func passArray(items [size]int) {
	items[0] = 100
	report("inside passArray")
}

func passSlice(items []int) {
	report("inside passSlice")
}

func report(msg string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("[%s]\n", msg)
	fmt.Printf("\t> Memory Usage: %v KB\n", m.Alloc/1024)
}
