// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

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
