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
	nums := []int{1, 2, 3}
	s.Show("nums", nums)

	_ = append(nums, 4)
	s.Show("nums", nums)

	nums = append(nums, 4)
	s.Show("nums", nums)

	nums = append(nums, 9)
	s.Show("nums", nums)

	nums = append(nums, 4)
	s.Show("nums", nums)

	// or:
	// nums = append(nums, 4, 9)
	// s.Show("nums", nums)

	nums = []int{1, 2, 3}
	tens := []int{12, 13}

	nums = append(nums, tens...)
	s.Show("nums", nums)
}
