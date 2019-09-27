// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

type toy struct {
	title string
	price money
}

func (t *toy) print() {
	fmt.Printf("%-15s: %s\n", t.title, t.price.string())
}

func (t *toy) discount(ratio float64) {
	t.price *= money(1 - ratio)
}
