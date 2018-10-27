// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"os"
)

func main() {
	// WARNING: This program will error
	//          if you don't pass your name and lastname

	name, lastname := os.Args[1], os.Args[2]

	msg := "Your name is %s and your lastname is %s.\n"
	fmt.Printf(msg, name, lastname)
}
