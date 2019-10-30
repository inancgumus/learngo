// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

func main() {
	// constants repeat the previous type and expression
	const (
		min int = 1
		max     // int = 1
	)

	fmt.Println(min, max)

	// print the types of min and max
	fmt.Printf("%T %T\n", min, max)
}
