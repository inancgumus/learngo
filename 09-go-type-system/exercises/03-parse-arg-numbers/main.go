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

// ---------------------------------------------------------
// EXERCISE: Parse Arg Numbers
//
//  Use strconv.ParseInt function to get int8, int16, and
//  int32, and int64 values from command-line.
//
// HINT
//  The third argument to ParseInt function represents
//  the bitsize.
//
//  So, giving it 8 returns an int8 convertable value;
//  whereas 16 returns an int16 convertable value.
//
//  Please explore the documentation of ParseInt function
//  and learn how it works.
//
// EXPECTED OUTPUT
//   When runned like this:
//     go run main.go 50 25000 2000000 50000000000000000 00000100
//
//   It should return this:
//     int8 value is : 50
//     int16 value is: 25000
//     int32 value is: 2000000
//     int64 value is: 50000000000000000
//     00000100 is: 4
// ---------------------------------------------------------

func main() {
	// --------------------------------------
	// EXAMPLE:
	// --------------------------------------
	// How to get an int8 from command-line:
	// First argument should be a value of -128 to 127 range
	//
	// Second argument: 10 means decimal number
	// Third argument : 8 means 8-bits (int8)
	val, _ := strconv.ParseInt(os.Args[1], 10, 8)

	// Now the val variable is int64 because ParseInt
	// returns an int64. But, since I passed 8 as its third
	// argument, it returns int8 convertable value.
	//
	// Try running the program with a value of -128 to 127
	// Running it beyond that range will result in
	// either -128 or 127.
	fmt.Println("int8 value is:", int8(val))

	// --------------------------------------
	// NOW IT'S YOUR TURN!
	// --------------------------------------

	// 1. Get an int16 value using ParseInt and convert it and print it

	// 2. Get an int32 value using ParseInt and convert it and print it

	// 3. Get an int64 value using ParseInt and convert it and print it

	// 4. Get an int8 value using ParseInt and convert it and print it
	//    But this time, get the value in bits.
	//
	//    For example : 00000100
	//    Should print: 4
}
