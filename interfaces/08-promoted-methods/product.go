// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

type product struct {
	title string
	price money
}

func (p *product) print() {
	fmt.Printf("%-15s: %s\n", p.title, p.price.string())
}

func (p *product) discount(ratio float64) {
	p.price *= money(1 - ratio)
}
