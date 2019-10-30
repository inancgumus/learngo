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

// be consistent:
// if another method in the same type has a pointer receiver,
// use pointer receivers on the other methods as well.
func (g *game) print() {
	fmt.Printf("%-15s: $%.2f\n", g.title, g.price)
}

// + discount gets a copy of `*game`.
// + discount can update the original `game` through the game pointer.
// + it's better to use the same receiver type: `*game` for all methods.
func (g *game) discount(ratio float64) {
	g.price *= (1 - ratio)
}

// PREVIOUS CODE:

// + `g` is a copy: `discount` cannot change the original `g`.
// func (g *game) discount(ratio float64) {
// 	g.price *= (1 - ratio)
// }
