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
)

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

	// ------------------------------------------------
	// How to calculate bits?
	// ------------------------------------------------

	// note: ^ means power of.
	//       for example: 2^2 means 2 * 2     = 4
	//                or: 2^3 means 2 * 2 * 2 = 8

	// 0 0 0 0 0 0 1 0 is equal to 2 because:
	// ^ ^ ^ ^ ^ ^ ^ ^
	// | | | | | | | |
	// | | | | | | | +-> 2^0 * 0 = 0
	// | | | | | | +---> 2^1 * 1 = 2
	// | | | | | +-----> 2^2 * 0 = 0
	// | | | | +-------> 2^3 * 0 = 0
	// | | | +---------> 2^4 * 0 = 0
	// | | +-----------> 2^5 * 0 = 0
	// | +-------------> 2^6 * 0 = 0
	// +---------------> 2^7 * 0 = 0
	//                   ------------+
	//                             2

	i, _ := strconv.ParseInt("00000010", 2, 64)
	fmt.Println(i)

	// 0 0 0 1 0 1 1 0 is equal to 22 because:
	// ^ ^ ^ ^ ^ ^ ^ ^
	// | | | | | | | |
	// | | | | | | | +-> 2^0 * 0 = 0
	// | | | | | | +---> 2^1 * 1 = 2
	// | | | | | +-----> 2^2 * 1 = 4
	// | | | | +-------> 2^3 * 0 = 0
	// | | | +---------> 2^4 * 1 = 16
	// | | +-----------> 2^5 * 0 = 0
	// | +-------------> 2^6 * 0 = 0
	// +---------------> 2^7 * 0 = 0
	//                   ------------+
	//                             22

	i, _ = strconv.ParseInt("00010110", 2, 64)
	fmt.Println(i)
}
