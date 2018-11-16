// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	// you cannot use variables
	// when setting the length of an array

	// length := 3

	// but, you can use constants and constant expressions
	const length = 3 * 2

	ages := [3 * 2]int{15, 30, 25}

	for _, age := range ages {
		fmt.Println(age)
	}
}

// EXERCISE:
//   Try to put the `length` constant
//   in place of `3 * 2` above.

// EXERCISE:
//   Try to put the `length` variable
//   in place of `3 * 2` above.
//
//   And observe the error.
//
//   To do that, you need comment-out
//   the length constant first.
