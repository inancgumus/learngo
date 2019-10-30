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
	// var counter int
	// var factor float64
	var (
		counter int
		factor  float64
	)

	// counter = counter + 1
	counter++
	fmt.Println(counter)

	// counter = counter - 1
	counter--
	fmt.Println(counter)

	// counter = counter + 5
	counter += 5
	fmt.Println(counter)

	// counter = counter * 10
	counter *= 10
	fmt.Println(counter)

	// counter = counter / 2.0
	counter /= 2.0
	fmt.Println(counter)

	factor += float64(counter)
	fmt.Println(counter)

	var bigCounter int64
	counter += int(bigCounter)
	fmt.Println(counter)

	fmt.Println(
		"hello" + ", " + "how" + " " + "are" + " " + "today?",
	)

	// you can combine raw string and string literals
	fmt.Println(
		`hello` + `, ` + `how` + ` ` + `are` + ` ` + "today?",
	)

	// ------------------------------------------
	// Converting non-string values into string
	// ------------------------------------------

	eq := "1 + 2 = "
	sum := 1 + 2

	// invalid op
	// string concat op can only be used with strings
	// fmt.Println(eq + sum)

	// you need to convert it using strconv.Itoa
	// Itoa = Integer to ASCII

	fmt.Println(eq + strconv.Itoa(sum))
}
