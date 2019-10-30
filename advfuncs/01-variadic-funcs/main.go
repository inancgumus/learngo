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
	nums := []int{2, 3, 7}
	fmt.Printf("nums            : %d\n", nums)

	n := avgNoVariadic(nums)
	fmt.Printf("avgNoVariadic   : %d\n", n)

	n = avg(2, 3, 7)
	fmt.Printf("avg(2, 3, 7)    : %d\n", n)

	n = avg(2, 3, 7, 8)
	fmt.Printf("avg(2, 3, 7, 8) : %d\n", n)

	// use ... to pass a slice
	n = avg(nums...)
	fmt.Printf("avg(nums...)    : %d\n", n)

	// uses the existing slice
	double(nums...)
	fmt.Printf("double(nums...) : %d\n", nums)

	// creates a new slice
	double(4, 6, 14)
	fmt.Printf("double(4, 6, 14): %d\n", nums)

	// creates a nil slice
	fmt.Printf("\nmain.nums       : %p\n", nums)
	investigate("passes main.nums", nums...)
	investigate("passes main.nums", nums...)
	investigate("passes args", 4, 6, 14)
	investigate("passes args", 4, 6, 14)
	investigate("no args")
}

func investigate(msg string, nums ...int) {
	fmt.Printf("investigate.nums: %12p  ->  %s\n", nums, msg)

	if len(nums) > 0 {
		fmt.Printf("\tfirst element: %d\n", nums[0])
	}
}

func double(nums ...int) {
	for i := range nums {
		nums[i] *= 2
	}
}

func avg(nums ...int) int {
	return sum(nums) / len(nums)
}

func avgNoVariadic(nums []int) int {
	return sum(nums) / len(nums)
}

func sum(nums []int) (total int) {
	for _, n := range nums {
		total += n
	}
	return
}
