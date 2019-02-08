// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Fix the backing array problem
//
//  You receive numbers from an API. After you're done working
//  with it, the API needs to continue using those numbers.
//
//  But your program changes the numbers (changes the API's slice).
//
//  Fix the program so that your program doesn't modify
//  the original numbers.
//
//
// RESTRICTION
//
//  Fix the problem only in the designated area of the code below.
//
//
// EXPECTED OUTPUT
//
//  Mine         : [-50 -100 -150]
//  Original nums: [56 89 15]
//
//  Note: Original nums may vary (they're random)
//        But your slice should look like the above (mine slice)
//
//        Yes, it should output only three numbers for the both slices!
//
// ---------------------------------------------------------

func main() {
	// API returns random numbers in an int slice
	rand.Seed(time.Now().UnixNano())
	nums := rand.Perm(100)

	// ----------------------------------------
	// RESTRICTIONS — ONLY ADD YOUR CODE HERE
	//
	mine := nums
	//
	// ----------------------------------------

	mine[0], mine[1], mine[2] = -50, -100, -150
	fmt.Println("Mine         :", mine)
	fmt.Println("Original nums:", nums[:3])
}
