// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

type filterFunc func(int) bool

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("••• FUNC VALUES •••")
	fmt.Printf("evens      : %d\n", filter(isEven, nums...))
	fmt.Printf("odds       : %d\n", filter(isOdd, nums...))
	fmt.Printf("> 3        : %d\n", filter(greaterThan3, nums...))
	fmt.Printf("> 6        : %d\n", filter(greaterThan6, nums...))

	// ========================================================================

	fmt.Println("\n••• CLOSURES •••")

	var min int
	greaterThan := func(n int) bool {
		return n > min
	}

	min = 3
	fmt.Printf("> 3        : %d\n", filter(greaterThan, nums...))

	min = 6
	fmt.Printf("> 6        : %d\n", filter(greaterThan, nums...))

	// min = 1
	// fmt.Printf("> 1        : %d\n", filter(greaterThan, nums...))
	// min = 2
	// fmt.Printf("> 2        : %d\n", filter(greaterThan, nums...))
	// min = 3
	// fmt.Printf("> 3        : %d\n", filter(greaterThan, nums...))

	var filterers []filterFunc
	for i := 1; i <= 3; i++ {
		current := i

		filterers = append(filterers, func(n int) bool {
			min = current
			return greaterThan(n)
		})
	}

	printer(filterers, nums...)
}

func printer(filterers []filterFunc, nums ...int) {
	for i, filterer := range filterers {
		fmt.Printf("> %d        : %d\n", i+1, filter(filterer, nums...))
	}
}

func greaterThan6(n int) bool { return n > 6 }
func greaterThan3(n int) bool { return n > 3 }
func isEven(n int) bool       { return n%2 == 0 }
func isOdd(m int) bool        { return m%2 == 1 }

func filter(f filterFunc, nums ...int) (filtered []int) {
	for _, n := range nums {
		if f(n) {
			filtered = append(filtered, n)
		}
	}
	return
}
