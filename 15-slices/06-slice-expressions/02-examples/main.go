// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	// 1. slicing creates a new slice
	nums := []int{1, 3, 5, 2, 4, 8}
	odds := nums[:3]
	evens := nums[3:]

	// 2. backing array is shared
	nums[1], nums[3] = 9, 6
	s.Show("nums", nums)
	s.Show("odds : nums[:3]", odds)
	s.Show("evens: nums[3:]", evens)

	// 3. you can create a new slice from an array
	heyArr := [3]byte{'h', 'e', 'y'}
	hey, he := heyArr[:], heyArr[:2]

	// 4. sliced array will be shared among the slices
	heyArr[0] = 'm'
	s.Show("hey := heyArr[:]", hey)
	s.Show("he  := heyArr[:2]", he)

	// 5. index expression returns a value
	//    while a slice expression returns a slice
	s.Show("nums[0]", nums[0])
	s.Show("nums[0:1]", nums[0:1])
	fmt.Printf("nums[0]  : %T\n", nums[0])   // just int
	fmt.Printf("nums[0:1]: %T\n", nums[0:1]) // []int slice

	// 6. extending a slice up to its capacity
	first := nums[0:1]
	s.Show("first := nums[0:1]", first)
	s.Show("first[0:2]", first[0:2])
	s.Show("first[0:3]", first[0:3])
	s.Show("first[0:4]", first[0:4])
	s.Show("first[0:5]", first[0:5])
	s.Show("first[0:6]", first[0:6])
	// s.Show("first[0:7]", first[0:7]) // <- you can't

	first = append(first[0:6], 9)
	s.Show("first: with a new backing array", first)
}
