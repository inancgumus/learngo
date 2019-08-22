// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"

	"github.com/inancgumus/learngo/16-slices/exercises/23-limit-the-backing-array-sharing/api"
)

// ---------------------------------------------------------
// EXERCISE: Limit the backing array sharing
//
//  In this exercise: API means the api package. It's in the
//  api folder.
//
//  `Read` function of the api package returns a portion of
//  a slice. The main function [main()] below uses the
//  Read function.
//
// `main()` appends to the slice but doing so changes the
//  backing array of the api package's temps slice as well.
//  We don't want that.
//
//  Only allow `main()` to change the part of the slice
//  it receives from the Read function.
//
//
// NOTE
//
//   You need to import the api package.
//
//
// CURRENT
//
//                           | |
//                           v v
//   api.temps     : [5 10 3 1 3 80 90]
//   main.temps    : [5 10 3 1 3]
//                           ^ ^ append changes the api package's
//                               temps slice's backing array.
//
//
//
// EXPECTED
//
//   The corrected api package does not allow the `main()` change
//   the api package's temps slice's backing array.
//                           |  |
//                           v  v
//   api.temps     : [5 10 3 25 45 80 90]
//   main.temps    : [5 10 3 1 3]
//
// ---------------------------------------------------------

func main() {
	// DO NOT CHANGE ANYTHING IN THIS CODE.

	// get the first three elements from api.temps
	slice := api.Read(0, 3)

	// append changes the api package's temps slice's
	// backing array as well.
	slice = append(slice, []int{1, 3}...)

	fmt.Println("api.temps     :", api.All())
	fmt.Println("main.slice    :", slice)
}
