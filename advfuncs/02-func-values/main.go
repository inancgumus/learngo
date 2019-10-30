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
	signatures()
	funcValues()
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

func funcValues() {
	fmt.Println("\n••• FUNC VALUES (VARS) •••")

	// var value func(int) bool
	var value filterFunc
	fmt.Printf("value nil? : %t\n", value == nil)

	value = isEven
	fmt.Printf("value(2)   : %t\n", value(2))
	fmt.Printf("isEven(2)  : %t\n", isEven(2))

	value = isOdd
	fmt.Printf("value(1)   : %t\n", value(1))
	fmt.Printf("isOdd(1)   : %t\n", isOdd(1))

	fmt.Printf("value nil? : %t\n", value == nil)
	fmt.Printf("value      : %p\n", value)
	fmt.Printf("isEven     : %p\n", isEven)
	fmt.Printf("isOdd      : %p\n", isOdd)
}
