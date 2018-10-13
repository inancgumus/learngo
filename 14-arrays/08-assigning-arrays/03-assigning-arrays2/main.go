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
	b := [3]int{6, 9, 3}

	b = a

	fmt.Println("b =>", b)

	b[0] = 100
	fmt.Println("a =>", a)
	fmt.Println("b =>", b)
}
