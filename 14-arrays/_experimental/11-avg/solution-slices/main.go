// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	var (
		sum  float64
		nums []float64
	)

	for _, v := range args {
		n, err := strconv.ParseFloat(v, 64)
		if err != nil {
			continue
		}

		nums = append(nums, n)
		sum += n
	}

	fmt.Println("Your numbers:", nums)
	fmt.Println("Average:", sum/float64(len(nums)))
}
