// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	var (
		nums   []int
		oldCap float64
	)

	for len(nums) < 10e6 {
		c := float64(cap(nums))

		if c == 0 || c != oldCap {
			fmt.Printf("len:%-15d cap:%-15g growth:%-15.2f\n",
				len(nums), c, c/oldCap)
		}
		oldCap = c

		nums = append(nums, 1)
	}
}
