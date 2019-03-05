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
//  You've created an API that returns population counts in a country.
//  To do that, you return an int slice up to some portion of it.
//
//  There is a program that uses your API but it appends to the slice
//  that your API returns. Doing so, overwrites your API's slice's
//  backing array as well.
//
//  Change your API so that it prevents the overwriting when
//  the client code wants to append to the returned slice from your
//  API.
//
//
// STEPS
//
//   1. Open the code inside the `api/api.go` folder
//
//   2. Fix the code there (not here — but run this code after)
//
//
// CURRENT OUTPUT
//
//       The following program overwrites the elements incorrectly
//          You need to change your API to prevent this behavior
//                           ^ ^
//                           | |
//   API's readings: [5 10 3 1 3 80 90]
//   Your readings : [5 10 3 1 3]
//
//
// EXPECTED OUTPUT
//
//   Now the program cannot change the API's original backing array
//   (beyond the returned capacity) (so the api now owns the control)
//                           ^  ^
//                           |  |
//   API's readings: [5 10 3 25 45 80 90]
//   Your readings : [5 10 3 1 3]
//
// ---------------------------------------------------------

func main() {
	// DO NOT TOUCH THE FOLLOWING CODE
	// THIS IS THE CLIENT PROGRAM THAT USES YOUR API
	// YOU CANNOT CONTROL IT! :)

	// reads the first three temperatures
	temps := api.Read(0, 3)

	// appends two new temperature readings
	temps = append(temps, []int{1, 3}...)

	// prints the current temperatures
	fmt.Println("API's readings:", api.All())
	fmt.Println("Your readings :", temps)
}
