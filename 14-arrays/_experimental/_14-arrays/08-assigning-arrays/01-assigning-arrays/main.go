// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	a := [3]int{5, 4, 7}

	// this assignment will create a new array value
	// and then it will be stored in the b variable
	b := a

	fmt.Println("a =>", a)
	fmt.Println("b =>", b)

	// they're equal, they're clones
	fmt.Println("a==b?", a == b)
}
