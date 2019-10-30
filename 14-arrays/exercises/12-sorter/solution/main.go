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
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	switch l := len(args); {
	case l == 0:
		fmt.Println("Please give me up to 5 numbers.")
		return
	case l > 5:
		fmt.Println("Sorry. Go arrays are fixed.",
			"So, for now, I'm only supporting sorting 5 numbers...")
		return
	}

	var nums [5]float64

	// fill the array with the numbers
	for i, v := range args {
		n, err := strconv.ParseFloat(v, 64)
		if err != nil {
			// skip if it's not a valid number
			continue
		}

		nums[i] = n
	}

	/*
		check whether it's the last element or not:
		  i < len(nums)-1

		check whether the next number is greater than the current one, if so, swap it:
		  v > nums[i+1]
	*/
	for range nums {
		for i, v := range nums {
			if i < len(nums)-1 && v > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}

	fmt.Println(nums)
}
