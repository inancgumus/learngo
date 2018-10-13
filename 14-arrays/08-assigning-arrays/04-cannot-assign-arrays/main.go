// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

func main() {
	// you can't assign the array 'a' to the array 'b'
	// a's type is [2]int whereas b's type is [3]int

	// when two values are not comparable,
	// then they're not assignable either

	// uncomment and observe the error

	// a := [2]int{5, 4}
	// b := [3]int{6, 9, 3}

	// b = a
}
