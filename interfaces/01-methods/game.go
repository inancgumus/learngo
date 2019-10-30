// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

type game struct {
	title string
	price float64
}

func (g game) print() {
	fmt.Printf("%-15s: $%.2f\n", g.title, g.price)
}

// PREVIOUS CODE:

// ----------------------------------------------------------------------------
// you can use same method name among different types.
// you don't need to type `printGame`, it's just: `print`.
//
// func (g game) printGame() {
// 	fmt.Printf("%-15s: $%.2f\n", g.title, g.price)
// }

// ----------------------------------------------------------------------------
// you cannot use the same function name within the same package.
//
// func printGame(g game) {
// 	fmt.Printf("%-15s: $%.2f\n", g.title, g.price)
// }
