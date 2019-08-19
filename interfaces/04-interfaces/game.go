// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

type game struct {
	title string
	price float64
}

func (g *game) print() {
	fmt.Printf("%-15s: $%.2f\n", g.title, g.price)
}

func (g *game) discount(ratio float64) {
	g.price *= (1 - ratio)
}

// TODO: NEW
func (g *game) total() float64 {
	return g.price
}
