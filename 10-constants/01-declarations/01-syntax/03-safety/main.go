// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	// Go can't catch the same error at runtime
	// When you run this, there will be an error:
	//
	// panic: runtime error: integer divide by zero
	n, m := 1, 0
	fmt.Println(n / m)

	// Go will detect the division by zero error
	// at compile-time
	//
	// const n int = 1
	// const m int = 0
	// fmt.Println(n / m)
}
