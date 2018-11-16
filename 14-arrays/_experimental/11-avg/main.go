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
		nums [5]float64
	)

	for i, v := range args {
		n, err := strconv.ParseFloat(v, 64)
		if err != nil {
			continue
		}

		nums[i] = n
		sum += n
	}

	fmt.Println("Your numbers   :", nums)
	fmt.Println("Average        :", sum/float64(len(nums)))
}

// EXERCISE
// Average calculator has a flaw.
// It divides the numbers by the length of the array.
// This results in wrong calculatio.
//
// For example:
//
// When you run it like this:
//   go run main.go 1 5
// It tells you that the average number is:
//  1.2
// Whereas, actually, it should be 3.
//
// Fix this error.
// So that, it will output 3 instead of 1.2
//
// Do not change the length of the array.
