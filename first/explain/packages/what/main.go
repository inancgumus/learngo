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
	fmt.Println("Hello!")

	// You can access functions from other files
	// which are in the same package

	// Here, `main()` can access `bye()` and `hey()`

	// It's because bye.go, hey.go and main.go
	//   are in the main package.

	bye()
	hey()
}
