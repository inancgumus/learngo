// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	s "github.com/inancgumus/prettyslice"
)

func main() {
	s.PrintBacking = true

	// #1: a nil slice has no backing array
	var nums []int
	s.Show("no backing array", nums)

	// #2: creates a new backing array
	nums = append(nums, 1, 3)
	s.Show("allocates", nums)

	// #3: creates a new backing array
	nums = append(nums, 2)
	s.Show("free capacity", nums)

	// #4: uses the same backing array
	nums = append(nums, 4)
	s.Show("no allocation", nums)

	// GOAL: append new odd numbers in the middle
	// [1 3 2 4] -> [1 3 7 9 2 4]

	// #5: [1 3 2 4] -> [1 3 2 4 2 4]
	nums = append(nums, nums[2:]...)
	s.Show("nums <- nums[2:]", nums)

	// #6: overwrites: [1 3 2 4 2 4] -> [1 3 7 9]
	nums = append(nums[:2], 7, 9)
	s.Show("nums[:2] <- 7, 9", nums)

	// #7: [1 3 7 9] -> [1 3 7 9 2 4]
	nums = nums[:6]
	s.Show("nums: extend", nums)
}

// don't mind about these options
// they're just for printing the slices nicely
func init() {
	s.MaxPerLine = 10
	s.Width = 45
}
