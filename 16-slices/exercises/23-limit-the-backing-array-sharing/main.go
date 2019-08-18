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
//  You've created a package and saved the temperatures in an int slice.
//  (find it in the api folder).
//
//  When `Read()` function is called, you return the desired elements.
//
//  Another programmer has decided to use your package.
//  She calls: `api.Read(0, 3)` and gets a slice with a length of 3.
//
//  However: She appends to the slice, so she practically changes
//  the elements in the api package's `temps` slice (the original one).
//
//  Only allow her to change the first three elements.
//  Prevent her from changing the rest of the elements of
//  `api.temps` slice.
//
//
// STEPS
//
//   You only need to change the code inside the `api/api.go` folder
//
//
// CURRENT
//
//                           | |
//                           v v
//   api.temps     : [5 10 3 1 3 80 90]
//   main.temps    : [5 10 3 1 3]
//                           ^ ^ append changes the api package's
//                               temps slice: [1 3]
//
//
//
// EXPECTED
//
//   Now the program cannot change the API's original backing array.
//                           |  |
//                           v  v
//   api.temps     : [5 10 3 25 45 80 90]
//   main.temps    : [5 10 3 1 3]
//
// ---------------------------------------------------------

func main() {
	// DO NOT TOUCH THE FOLLOWING CODE
	// THIS IS THE CLIENT PROGRAM THAT USES YOUR API
	// YOU CANNOT CONTROL IT! :)

	// get the first three elements from api.temps
	temps := api.Read(0, 3)

	// changes the api.temps's backing array.
	// you need to prevent this.
	temps = append(temps, []int{1, 3}...)

	fmt.Println("api.temps     :", api.All())
	fmt.Println("main.temps    :", temps)
}
