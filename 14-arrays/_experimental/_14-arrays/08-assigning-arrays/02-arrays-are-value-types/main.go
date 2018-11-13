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
	b := a

	// this will only change the 'a' array
	a[0] = 10

	fmt.Println("a =>", a)
	fmt.Println("b =>", b)

	// this will only change the 'b' array
	b[1] = 20

	fmt.Println("a =>", a)
	fmt.Println("b =>", b)
}
