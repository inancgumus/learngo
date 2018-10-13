// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

func main() {
	// -----
	// you cannot use incompatible values
	// than the array's element type
	//
	// uncomment the code parts below one by one
	// then observe the errors
	// -----

	// for example you can't use a string
	// ages := [3]int{15, 30, "hi"}

	// -----

	// or a float64, etc...
	// ages := [3]int{15, 30, 5.5}

	// -----

	// of course this is valid, since it's untyped
	// 5. becomes 5 (int)
	// ages := [3]int{15, 30, 5.}

	// -----

	// for _, age := range ages {
	// 	fmt.Println(age)
	// }
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
