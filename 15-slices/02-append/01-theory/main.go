// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	s "github.com/inancgumus/prettyslice"
)

func main() {
	nums := []int{1, 2, 3}
	s.Show("nums := []int{1, 2, 3}", nums)

	// -----------------------------------------------------
	nums = append(nums, 4)
	s.Show("append(nums, 4)", nums)

	// -----------------------------------------------------
	nums = append(nums, 9)
	s.Show("append(nums, 9)", nums)

	// -----------------------------------------------------
	// let's reset nums
	// and let's add multiple elements
	nums = []int{1, 2, 3}
	s.Show("nums = []int{1, 2, 3}", nums)

	nums = append(nums, 4, 9)
	s.Show("append(nums, 4, 9)", nums)

	// -----------------------------------------------------
	// let's reset nums again
	// let's add multiple elements using the ellipsis
	nums = []int{1, 2, 3}
	tens := []int{12, 13}

	s.Show("nums = []int{1, 2, 3}", nums)
	s.Show("tens := []int{12, 13}", tens)

	nums = append(nums, tens...)
	s.Show("append(nums, tens...)", nums)
}
