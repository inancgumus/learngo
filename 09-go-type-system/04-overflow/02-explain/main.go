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
	"math"
)

func main() {
	// go catches overflow at compile-time
	//
	// fmt.Println(int8(math.MaxInt8 + 1)) // overflows

	// but it cannot catch them in runtime
	var n int8 = math.MaxInt8

	// wrap arounds to its negative maximum
	fmt.Println("max int8     :", n)   // 127
	fmt.Println("max int8 + 1 :", n+1) // -128

	// wrap arounds to its positive maximum
	n = math.MinInt8
	fmt.Println("min int8     :", n)   // -128
	fmt.Println("min int8 - 1 :", n-1) // 127

	// wrap arounds to its maximum
	var un uint8

	fmt.Println("max uint8    :", un)   // 0
	fmt.Println("max uint8 - 1:", un-1) // 255

	// wrap around to its minimum
	un = math.MaxUint8
	fmt.Println("max uint8 + 1:", un+1) // 0

	// floats goes to infinity when overflowed
	f := float32(math.MaxFloat32)
	fmt.Println("max float    :", f*1.2)
}
