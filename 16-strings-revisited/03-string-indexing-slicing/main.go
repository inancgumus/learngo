// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
)

func main() {
	// GOALS:
	// 1- String value is immutable
	// 2- Indexing vs Slicing
	// 3- Using bytes for manipulating strings

	mood := "wonder"

	// 1- a string value is immutable (read-only)
	// mood[1] = 'a'

	// 2- Indexing vs Slicing

	// "wonder"
	//  ^ ^^^^
	//  | ||||
	// "wandering"

	//           "w"     + "a" + "nder"   + "ing"

	// wandering := mood[0] + "a" + mood[2:] + "ing"
	// fmt.Printf("mood[0]   : %T - %[1]v\n", mood[0])   // byte
	// fmt.Printf("mood[0:1] : %T - %[1]v\n", mood[0:1]) // string

	// wandering := mood[:1] + "a" + mood[2:] + "ing"
	fmt.Println(mood)
	// fmt.Println(wandering)

	// 3- converting creates a new byte slice (allocation)
	b := []byte(mood)
	b[1] = 'a'

	// b = append(b, 'i', 'n', 'g')
	// b = append(b, []byte{'i', 'n', 'g'})
	b = append(b, "ing"...)

	// starts copying from the first element
	copy(b, "listen")

	// starts copying from the "7th" element
	copy(b[6:], "ed.")

	fmt.Println(string(b))
}
