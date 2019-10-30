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
	"strconv"
	"strings"
)

func main() {
	data := "2 4 6 1 3 5"
	splitted := strings.Fields(data)

	var nums []int
	for _, s := range splitted {
		n, _ := strconv.Atoi(s)
		nums = append(nums, n)
	}

	// rest of the code for this exercise
	fmt.Println("nums        :", nums)

	evens, odds := nums[:3], nums[3:]

	fmt.Println("evens       :", evens)
	fmt.Println("odds        :", odds)

	fmt.Println("middle      :", nums[2:4])
	fmt.Println("first 2     :", nums[:2])
	fmt.Println("last 2      :", nums[len(nums)-2:])

	fmt.Println("evens last 1:", evens[len(evens)-1:])
	fmt.Println("odds last 2 :", odds[len(odds)-2:])
}
