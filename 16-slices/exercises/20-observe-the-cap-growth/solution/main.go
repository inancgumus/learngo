// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

func main() {
	var (
		nums   []int
		oldCap float64
	)

	// loop 10 million times
	for len(nums) < 1e7 {
		// get the capacity
		c := float64(cap(nums))

		// only print when the capacity changes
		if c == 0 || c != oldCap {
			// print also the growth ratio: c/oldCap
			fmt.Printf("len:%-15d cap:%-15g growth:%-15.2f\n",
				len(nums), c, c/oldCap)
		}

		// keep track of the previous capacity
		oldCap = c

		// append an arbitrary element to the slice
		nums = append(nums, 1)
	}
}
