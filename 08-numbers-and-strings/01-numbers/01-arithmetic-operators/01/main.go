// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	// when an integer and a float value used together
	// in an expression, the result always becomes
	// a float value
	fmt.Println(8 * -4.0) // -32.0 not -32

	// two integer values result in an integer value
	fmt.Println(-4 / 2)

	// remainder operator
	// it can only used with integers
	fmt.Println(5 % 2)
	// fmt.Println(5.0 % 2) // wrong

	// addition operators
	fmt.Println(1 + 2.5)
	fmt.Println(2 - 3)

	// negation operator
	fmt.Println(-(-2))
	fmt.Println(- -2) // this also works
}
