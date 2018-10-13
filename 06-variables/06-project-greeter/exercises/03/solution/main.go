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
	fmt.Println(os.Args[1])

	// BONUS SOLUTION:
	fmt.Println("Hello", os.Args[1])
	fmt.Println("How are you?")
}
