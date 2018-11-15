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
		sum   float64
		nums  [5]float64
		total float64
	)

	for i, v := range args {
		n, err := strconv.ParseFloat(v, 64)
		if err != nil {
			continue
		}

		total++
		nums[i] = n
		sum += n
	}

	fmt.Println("Your numbers:", nums)
	fmt.Printf("(Only %g of them were valid)\n", total)
	fmt.Println("Average:", sum/total)
}
