// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"strings"
)

func main() {
	//
	// DAILY REQUESTS DATA
	//
	reqs := []int{
		500, 600, 250,
		200, 400, 50,
		900, 800, 600,
		750, 250, 100,
		100, 150,
	}

	//
	// There are 3 requests per day
	//
	const N = 3

	//
	// Allocate the slice efficiently with the exact size needed
	//
	l := len(reqs)
	size := l / N
	if l%N != 0 {
		size++
	}
	daily := make([][]int, 0, size)

	//
	// Group the `reqs` per day into a slice named: `daily`
	//
	for N < len(reqs) {
		daily = append(daily, reqs[:N]) // add the current batch of nums to the `groups`
		reqs = reqs[N:]                 // move the slice pointer for the next batch
	}

	//
	// Add the residual elements to the group (len(reqs) % N)
	//
	daily = append(daily, reqs)

	//
	// Print the header:
	//
	fmt.Printf("%-10s%-10s\n", "Day", "Requests")
	fmt.Println(strings.Repeat("=", 20))

	//
	// Print the data per day along with the totals:
	//
	var grand int

	for i, d := range daily {
		var sum int

		for _, q := range d {
			sum += q
			fmt.Printf("%-10d%-10d\n", i+1, q)
		}

		fmt.Println(strings.Repeat("-", 20))
		fmt.Printf("%10s%-10d\n\n", "", sum)

		grand += sum
	}

	fmt.Printf("%10s%-10d\n", "", grand)
}
