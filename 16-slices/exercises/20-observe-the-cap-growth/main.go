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
	"math/rand"
)

// ---------------------------------------------------------
// EXERCISE: Observe the capacity growth
//
//  Write a program that appends elements to a slice
//  10 million times in a loop. Observe how the capacity of
//  the slice changes.
//
//
// STEPS
//
//  1. Create a nil slice
//
//  2. Loop 1e7 times
//
//  3. On each iteration: Append an element to the slice
//
//  4. Print the length and capacity of the slice "only"
//     when its capacity changes.
//
//  BONUS: Print also the growth rate of the capacity.
//
//
// EXPECTED OUTPUT
//
//  len:0               cap:0               growth:NaN
//  len:1               cap:1               growth:+Inf
//  len:2               cap:2               growth:2.00
//  ... and so on.
//
// ---------------------------------------------------------

func main() {
	//  1. Create a nil slice
	var my_slice []int
	var initial_cap float64 = 0.
	fmt.Printf("len:%d              cap:%d              growth:%.2f\n",
		len(my_slice), cap(my_slice), (float64(cap(my_slice)))/initial_cap)
	for i := 0; i < 1e7; i++ {
		my_slice = append(my_slice, rand.Intn(1e7+1))
		if initial_cap != float64(cap(my_slice)) {
			fmt.Printf("len:%d              cap:%d              growth:%.2f\n",
				len(my_slice), cap(my_slice), (float64(cap(my_slice)))/initial_cap)
			initial_cap = float64(cap(my_slice))

		}
	}

}
