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
	"strings"
)

// ---------------------------------------------------------
// CHALLENGE
//  Add error handling to the feet to meters program.
//
// EXPECTED OUTPUT
//  go run main.go hello
//    error: 'hello' is not a number.
//
//  go run main.go what
//    error: 'what' is not a number.
//
//  go run main.go 100
//    100 feet is 30.48 meters.
// ---------------------------------------------------------

const usage = `
Feet to Meters
--------------
This program converts feet into meters.

Usage:
feet [feetsToConvert]`

func main() {
	if len(os.Args) < 2 {
		fmt.Println(strings.TrimSpace(usage))
		return
	}

	arg := os.Args[1]

	feet, _ := strconv.ParseFloat(arg, 64)
	meters := feet * 0.3048

	fmt.Printf("%g feet is %g meters.\n", feet, meters)
}
