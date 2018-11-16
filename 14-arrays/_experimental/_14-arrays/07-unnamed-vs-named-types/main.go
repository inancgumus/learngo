// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	type (
		A [3]int
		B [3]int
	)

	x := A{1, 2, 3}
	y := B{1, 2, 3}
	z := [3]int{1, 2, 3}

	fmt.Printf("x's type: %T\n", x) // named type  : main.A
	fmt.Printf("y's type: %T\n", y) // named type  : main.B
	fmt.Printf("z's type: %T\n", z) // unnamed type: [3]int

	// cannot compare different named types
	// fmt.Println("x==y?", x == y)

	// can convert between identical types
	fmt.Println("x==y?", x == A(y))
	fmt.Println("x==y?", B(x) == y)

	fmt.Printf("x's type   : %T\n", x)
	fmt.Printf("A(y)'s type: %T\n", A(y))

	// can compare to unnamed types
	fmt.Println("x==z?", x == z)
}
