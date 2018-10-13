// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	n := 10

	// ALTERNATIVES:
	// n = n - 1
	// n -= 1

	// BETTER:
	n--

	fmt.Println(n)
}
