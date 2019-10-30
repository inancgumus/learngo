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
	{
		// its length is part of its type
		var nums [5]int
		fmt.Printf("nums array: %#v\n", nums)
	}

	{
		// its length is not part of its type
		var nums []int
		fmt.Printf("nums slice: %#v\n", nums)

		fmt.Printf("len(nums) : %d\n", len(nums))

		// won't work: the slice is nil.
		// fmt.Printf("nums[0]: %d\n", nums[0])
		// fmt.Printf("nums[1]: %d\n", nums[1])
	}
}
