// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

type printer interface {
	print()
}

type list []printer

func (l list) print() {
	if len(l) == 0 {
		fmt.Println("Sorry. We're waiting for delivery 🚚.")
		return
	}

	for _, it := range l {
		// fmt.Printf("(%-10T) --> ", it)

		it.print()

		// you cannot access to the discount method of the game type.
		// `it` is a printer not a game.
		// it.discount(.5)
	}
}

// PREVIOUS CODE:

// type list []*game

// func (l list) print() {
// 	if len(l) == 0 {
// 		fmt.Println("Sorry. Our store is closed. We're waiting for the delivery 🚚.")
// 		return
// 	}

// 	for _, it := range l {
// 		it.print()
// 	}
// }
