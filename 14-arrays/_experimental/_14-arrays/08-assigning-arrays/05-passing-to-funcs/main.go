// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	a := [3]int{10, 4, 7}

	// like the assignment passing an array to a function
	// creates a copy of the array
	fmt.Println("a (before) =>", a)
	incr(a)
	fmt.Println("a (after)  =>", a)
}

// inside the function
// the passed array 'a' is a new copy (a clone)
// it's not the original one
func incr(a [3]int) {
	// this only changes the copy of the passed array
	a[1]++

	// this prints the copied array not the original one
	fmt.Println("a (inside) =>", a)
}
