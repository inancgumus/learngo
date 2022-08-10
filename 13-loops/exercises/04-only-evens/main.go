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

// ---------------------------------------------------------
// EXERCISE: Only Evens
//
//  1. Extend the "Sum up to N" exercise
//  2. Sum only the even numbers
//
// RESTRICTIONS
//  Skip odd numbers using the `continue` statement
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------

func main() {
	args := os.Args
	min, minErr := strconv.Atoi(args[1])
	max, maxErr := strconv.Atoi(args[2])

	if minErr != nil || maxErr != nil {
		fmt.Printf("%s min or %s max is not a number", args[1], args[2])
		return
	}

	if len(args) <= 2 {
		fmt.Println("More arguments needed. cmd min max")
		return
	}
	if len(args) == 3 {
		var sum int
		for i := min; i <= max; i++ {
			if i%2 == 0 {
				sum += i
				fmt.Printf("%1d", i)
				if i != max {
					fmt.Printf("+")
				}
			}

		}
		fmt.Printf("=%d", sum)
	}
}
