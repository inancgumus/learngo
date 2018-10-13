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
	n, err := strconv.Atoi("42")

	if err == nil {
		fmt.Println("There was no error, n is", n)
	}
}
