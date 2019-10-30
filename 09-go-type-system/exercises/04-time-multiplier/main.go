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
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Time Multiplier
//
//  You should get an argument from the command-line and
//  you need to multiply the time duration value `t` with
//  the given argument.
//
//  1. Get an argument from the command-line
//  2. Convert it to int64 and store it in a variable
//  3. Multiply the `t` variable with that variable
//  4. Print the result
//
// HINT
//  You can use ParseInt to convert the command-line
//    argument into an int64 value.
//
//  You can skip the error value using a blank-identifier.
//
// EXPECTED OUTPUT
//
//  When runned like this:
//    go run main.go 2
//
//  It should print this:
//    3h0m0s
// ---------------------------------------------------------

func main() {
	// DONT TOUCH THIS
	// 1h30m means: 1 hour 30 minutes
	t, _ := time.ParseDuration("1h30m")

	// TYPE YOUR CODE HERE
	// ....

	// DONT TOUCH THIS
	fmt.Println(t)
}
