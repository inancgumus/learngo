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
	// without newline
	fmt.Println("hihi")

	// with newline:
	//   \n = escape sequence
	//   \  = escape character
	fmt.Println("hi\nhi")

	// escape characters:
	//   \\ = \
	//   \" = "
	fmt.Println("hi\\n\"hi\"")
}
