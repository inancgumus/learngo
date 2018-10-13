// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	// When argument to len is a constant, len can be used
	// while initializing a constant
	//
	// Here, "Hello" is a constant.

	const max int = len("Hello")

	fmt.Println(max)
}
