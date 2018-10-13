// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
)

func main() {
	// cannot step over any variable declarations
	// ERROR: "i" variable is declared after the jump
	//
	// goto loop

	var i int

loop:
	if i < 3 {
		fmt.Println("looping")
		i++
		goto loop
	}
	fmt.Println("done")
}
