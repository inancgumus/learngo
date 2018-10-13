// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	// %b verb prints bits

	// true false, on off, ...
	// 1 bit can encode 2 different state: 0 or 1
	fmt.Printf("%b\n", 0)
	fmt.Printf("%b\n", 1)

	// 2 bits can encode 4 different states
	// 0 0
	// 0 1
	// 1 0
	// 1 1

	// %02b prints 2 bits with leading zeros if any

	fmt.Printf("%02b = %d\n", 0, 0)
	fmt.Printf("%02b = %d\n", 1, 1)
	fmt.Printf("%02b = %d\n", 2, 2)
	fmt.Printf("%02b = %d\n", 3, 3)

	// run this and analyze:
	// how 1 moves from right to the left

	// %08b prints 8 bits with leading zeros if any

	fmt.Printf("%08b = %d\n", 1, 1)
	fmt.Printf("%08b = %d\n", 2, 2)
	fmt.Printf("%08b = %d\n", 4, 4)
	fmt.Printf("%08b = %d\n", 8, 8)
	fmt.Printf("%08b = %d\n", 16, 16)
	fmt.Printf("%08b = %d\n", 32, 32)
	fmt.Printf("%08b = %d\n", 64, 64)
	fmt.Printf("%08b = %d\n", 128, 128)
}
