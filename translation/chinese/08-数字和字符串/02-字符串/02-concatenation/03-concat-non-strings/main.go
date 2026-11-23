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

	//

	// invalid op
	// eq = true + " " + false

	eq = strconv.FormatBool(true) +
		" " +
		strconv.FormatBool(false)

	fmt.Println(eq)
}
