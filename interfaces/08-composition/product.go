// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

type product struct {
	Title string
	Price money
}

func (p *product) discount(ratio float64) {
	p.Price *= money(1 - ratio)
}

func (p *product) sum() money {
	return p.Price
}

func (p *product) String() string {
	return fmt.Sprintf("%-15s: %s", p.Title, p.Price)
}
