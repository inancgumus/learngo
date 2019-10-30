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
	"strings"
	"unicode"
)

type filterFunc func(int) bool

func main() {
	signatures()

	fmt.Println("\n••• WITHOUT FUNC VALUES •••")
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Printf("evens      : %d\n", filterEvens(nums...))
	fmt.Printf("odds       : %d\n", filterOdds(nums...))

	fmt.Println("\n••• FUNC VALUES •••")
	fmt.Printf("evens      : %d\n", filter(isEven, nums...))
	fmt.Printf("odds       : %d\n", filter(isOdd, nums...))

	fmt.Println("\n••• MAPPING •••")
	fmt.Println(strings.Map(unpunct, "hello!!! HOW ARE YOU???? :))"))
	fmt.Println(strings.Map(unpunct, "TIME IS UP!"))
}

func unpunct(r rune) rune {
	if unicode.IsPunct(r) {
		return -1
	}
	return unicode.ToLower(r)
}

func filter(f filterFunc, nums ...int) (filtered []int) {
	for _, n := range nums {
		if f(n) {
			filtered = append(filtered, n)
		}
	}
	return
}

func filterOdds(nums ...int) (filtered []int) {
	for _, n := range nums {
		if isOdd(n) {
			filtered = append(filtered, n)
		}
	}
	return
}

func filterEvens(nums ...int) (filtered []int) {
	for _, n := range nums {
		if isEven(n) {
			filtered = append(filtered, n)
		}
	}
	return
}

func isEven(n int) bool {
	return n%2 == 0
}

func isOdd(m int) bool {
	return m%2 == 1
}

func signatures() {
	fmt.Println("••• FUNC SIGNATURES (TYPES) •••")
	fmt.Printf("isEven     : %T\n", isEven)
	fmt.Printf("isOdd      : %T\n", isOdd)
}
