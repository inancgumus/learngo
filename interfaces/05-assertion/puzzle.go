// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

type puzzle struct {
	title string
	price float64
}

func (p *puzzle) print() {
	fmt.Printf("%-15s: $%.2f \n", p.title, p.price)
}

func (p *puzzle) discount(ratio float64) {
	p.price *= (1 - ratio)
}
