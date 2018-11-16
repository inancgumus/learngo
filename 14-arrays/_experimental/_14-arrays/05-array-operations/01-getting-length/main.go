// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	ages := [2]int{15, 20}

	// i can directly use `len`
	fmt.Println(len(ages)) // 2

	// i can also assign its result to a variable
	l := len(ages)
	fmt.Println(l) // 2

	// let's print the length of a few arrays
	l = len([1]bool{true}) // 1
	fmt.Println(l)

	l = len([3]string{"lina", "james", "joe"}) // 3
	fmt.Println(l)

	// this array doesn't initialize any elements
	// but its length is still 5
	// whether the elements are initialized or not
	l = len([5]int{})
	fmt.Println(l) // 5
	fmt.Println([5]int{})
}
