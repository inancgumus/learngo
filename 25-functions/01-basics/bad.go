// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func showN() {
	if N == 0 {
		return
	}
	fmt.Printf("showN       : N is %d\n", N)
}

func incrN() {
	N++
}

// you cannot declare a function within the same package with the same name
// func incrN() {
// }
