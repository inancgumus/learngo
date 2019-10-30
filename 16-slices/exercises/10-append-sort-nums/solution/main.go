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
	"sort"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("provide a few numbers")
		return
	}

	var nums []int
	for _, s := range args {
		n, err := strconv.Atoi(s)
		if err != nil {
			continue
		}

		nums = append(nums, n)
	}

	sort.Ints(nums)
	fmt.Println(nums)
}
