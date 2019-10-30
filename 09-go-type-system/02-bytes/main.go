// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

func main() {
	// byte is an integer number with 8 bits (1 byte)
	var b byte

	// all bits are empty or 0
	// this is the minimum number a byte can represent
	b = 0
	fmt.Printf("%08b = %d\n", b, b)

	// all bits are full or 1
	// this is the maximum number a byte can represent
	b = 255
	fmt.Printf("%08b = %d\n", b, b)
}
