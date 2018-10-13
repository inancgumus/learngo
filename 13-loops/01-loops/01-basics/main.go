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
	var sum int

	// sum += 1
	// sum += 2
	// sum += 3
	// sum += 4
	// sum += 5

	for i := 1; i <= 1000; i++ {
		sum += i
	}

	fmt.Println(sum)
}
