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

	"github.com/inancgumus/learngo/16-slices/exercises/23-limit-the-backing-array-sharing/api"
)

// ---------------------------------------------------------
// EXERCISE: Limit the backing array sharing
//
//  GOAL
//
//    Limit the capacity of the slice that is returned
//    from the `Read` function. Read on for more details.
//
//
//  BEFORE YOU START
//
//    In this exercise: API means the api package. It's in the
//    api folder. You need to change the code in the `api/api.go`
//    to solve this exercise, and you need import the api
//    package.
//
//
//  WHAT IS THE PROBLEM?
//
//    `Read` function of the api package returns a portion of
//    its `temps` slice. Below, `main()` saves it to the
//    `received` slice.
//
//    `main()` appends to the `received` slice but doing so
//    also changes the backing array of the `temps` slice.
//    We don't want that.
//
//    `main()` can change the part of the `temps` slice
//    that is returned from the `Read()`, but it shouldn't
//    be able to change the elements in the rest of the
//    `temps`.
//
//
//  WHAT YOU NEED TO DO?
//
//    So you need to limit the capacity of the returned
//    slice somehow. Remember: `received` and `temps`
//    share the same backing array. So, appending to it
//    can overwrite the same backing array.
//
//
// CURRENT
//
//                           | |
//                           v v
//   api.temps     : [5 10 3 1 3 80 90]
//   main.received : [5 10 3 1 3]
//                           ^ ^ append changes the `temps`
//                               slice's backing array.
//
//
//
// EXPECTED
//
//   The corrected api package does not allow the `main()` to
//   change unreturned portion of the temps slice's backing array.
//                           |  |
//                           v  v
//   api.temps     : [5 10 3 25 45 80 90]
//   main.received : [5 10 3 1 3]
//
// ---------------------------------------------------------

func main() {
	// DO NOT CHANGE ANYTHING IN THIS CODE.

	// get the first three elements from api.temps
	received := api.Read(0, 3)

	// append changes the api package's temps slice's
	// backing array as well.
	received = append(received, []int{1, 3}...)

	fmt.Println("api.temps     :", api.All())
	fmt.Println("main.received :", received)
}
