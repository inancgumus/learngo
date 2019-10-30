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
	"time"
)

func main() {
	t, _ := time.ParseDuration("1h30m")

	// 1. get the first command-line argument
	// 2. convert it to int64
	multiplier, _ := strconv.ParseInt(os.Args[1], 10, 64)

	// 3. multiply the time duration with the given argument
	//
	//    converts the int64 value to time.Duration to be
	//    able to multiply it with the time.Duration value
	t *= time.Duration(multiplier)

	// 4. print it
	fmt.Println(t)
}
