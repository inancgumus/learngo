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
// EXERCISE: Find the Average
//
//  Your goal is to fill an array with numbers and find the average.
//
//  1. Get the numbers from the command-line.
//
//  2. Create an array and assign the given numbers to that array.
//
//  3. Print the given numbers and their average.
//
// RESTRICTION
//   + Maximum 5 numbers can be provided
//   + If one of the arguments are not a valid number, skip it
//
// EXPECTED OUTPUT
//   go run main.go
//     Please tell me numbers (maximum 5 numbers).
//
//   go run main.go 1 2 3 4 5 6
//     Please tell me numbers (maximum 5 numbers).
//
//   go run main.go 1 2 3 4 5
//     Your numbers: [1 2 3 4 5]
//     Average: 3
//
//   go run main.go 1 a 2 b 3
//     Your numbers: [1 0 2 0 3]
//     Average: 2
// ---------------------------------------------------------

func main() {
	input_args := os.Args[1:]
	switch n := len(input_args); {
	case n == 0:
		fmt.Println("Please tell me numbers (maximum 5 numbers).")
		return
	case n > 5:
		fmt.Print("Only 5 numbers can be provided.")
		return
	}
	numbers := [5]int{}
	sum := 0
	cnt := 0
	for ind, num := range input_args {
		val, err := strconv.Atoi(num)
		if err != nil {
			continue
		}
		numbers[ind] = val
		sum += val
		cnt += 1
	}
	fmt.Printf("Your numbers: %v\n", numbers)
	if cnt == 0 {
		fmt.Println("Average: 0")
		return
	}
	fmt.Printf("Average: %d\n", sum/cnt)
}
