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
	"unsafe"

	s "github.com/inancgumus/prettyslice"
)

// type collection [4]string // #1
type collection []string // #2

// go is pass by copy
// only the slice header is copied: 3 integer fields (24 bytes)
// think of passing an array with millions of elements instead.

func main() {
	// SliceHeader lives here:
	// https://golang.org/src/runtime/slice.go

	s.PrintElementAddr = true

	// #1
	data := collection{"slices", "are", "awesome", "period", "!!" /* #5 */}

	// data := collection{"slices", "are", "awesome", "period", "!!"}

	change(data) // #1

	s.Show("main's data", data)                           // #1
	fmt.Printf("main's data slice's header: %p\n", &data) // #3

	// ----------------------------------------------------------------------
	// #4
	array := [...]string{"slices", "are", "awesome", "period", "!!" /* #5 */}

	// array's size depends on its elements
	fmt.Printf("array's size: %d bytes.\n", unsafe.Sizeof(array))

	// slice's size is always fixed: 24 bytes (on a 64-bit system) — slice value = slice header
	fmt.Printf("slice's size: %d bytes.\n", unsafe.Sizeof(data))
}

// #1
// passed value will be copied in the function
func change(data collection) {
	// data is a new variable inside the function:
	// var data collection

	data[2] = "brilliant!"

	s.Show("change's data", data)
	fmt.Printf("change's data slice's header: %p\n", &data) // #3
}
