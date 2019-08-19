// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

// + you can attach methods to any concrete type.
// + rule: you need to declare a new type in the same package.
type list []game

func (l list) print() {
	// `list` acts like a `[]game`
	if len(l) == 0 {
		fmt.Println("Sorry. Our store is closed. We're waiting for the delivery 🚚.")
		return
	}

	for _, it := range l {
		it.print()
	}
}
