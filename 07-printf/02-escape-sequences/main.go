// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

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
