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
	// integer literal
	fmt.Println(
		-200, -100, 0, 50, 100, 100,
	)

	// float literal
	fmt.Println(
		-50.5, -20.5, -0., 1., 100.56, // ...
	)

	// bool constants
	fmt.Println(
		true, false,
	)

	// string literal - utf-8
	fmt.Println(
		"", // empty - prints just a space
		"hi",

		// unicode
		"nasılsın?",   // "how are you" in turkish
		"hi there 星!", // "hi there star!"
	)
}
