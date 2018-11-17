// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import s "github.com/inancgumus/prettyslice"

func main() {
	// #1: make preallocates an array
	nums := make([]int, 0, 10)
	nums = append(nums, 1, 2, 3)
	nums = append(nums, 4, 5)
	s.Show("nums", nums)

	nums = append(nums, 6)
	s.Show("nums: doesn't allocate", nums)

	// #2: prevent overwriting with make
	doubles := make([]int, len(nums) /*, 12 */)

	for i := range nums {
		doubles[i] = nums[i] * 2
	}

	s.Show("nums: after doubling", nums)
	s.Show("doubles", doubles)
}
