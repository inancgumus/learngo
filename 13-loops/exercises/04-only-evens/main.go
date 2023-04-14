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
	if len(os.Args) != 3 {
		fmt.Println("Please provide [start] [end].")
		return
	}
	start, err1 := strconv.Atoi(os.Args[1])
	end, err2 := strconv.Atoi(os.Args[2])

	if err1 != nil || err2 != nil || start >= end {
		fmt.Println("Please provide integers and make sure [start] < [end].")
		return
	}

	var total int
	if start%2 == 1 {
		start += 1
	}
	if end%2 == 1 {
		end -= 1
	}
	for i := start; i <= end; i += 2 {
		total += i
		fmt.Print(i)
		if i < end {
			fmt.Print(" + ")
		}
	}
	fmt.Printf(" = %d\n", total)

}
