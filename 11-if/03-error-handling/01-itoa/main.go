// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Itoa doesn't return any errors
	// So, you don't have to handle the errors for it

	s := strconv.Itoa(42)

	fmt.Println(s)
}
