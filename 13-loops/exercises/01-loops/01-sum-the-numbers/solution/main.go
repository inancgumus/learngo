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
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)
}
