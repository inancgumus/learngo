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

func main() {
	// FOR EXAMPLE:
	//
	// Run this program with 5 arguments like this:
	//
	//                int8  int16  int32    int64       as bits
	// go run main.go 50    25000  2000000  5000000000  00000100

	// first argument is an int8
	val, _ := strconv.ParseInt(os.Args[1], 10, 8)
	fmt.Println("int8 value is :", int8(val))

	// 2nd argument is an int16
	val, _ = strconv.ParseInt(os.Args[2], 10, 16)
	fmt.Println("int16 value is:", int16(val))

	// 3rd argument is an int32
	val, _ = strconv.ParseInt(os.Args[3], 10, 32)
	fmt.Println("int32 value is:", int32(val))

	// 4th argument is an int64
	// Remember ParseInt returns an int64
	val, _ = strconv.ParseInt(os.Args[4], 10, 64)
	fmt.Println("int64 value is:", val)

	// 5th argument is a number in bits. And its int8.
	// ParseInt(.., 2, ...) -> 2 means binary base
	val, _ = strconv.ParseInt(os.Args[5], 2, 8)
	fmt.Printf("%s is: %d\n", os.Args[5], int8(val))
}
