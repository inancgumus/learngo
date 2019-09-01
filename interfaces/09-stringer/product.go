// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

type product struct {
	title string
	price money
}

// product satisfies the fmt.Stringer
func (p *product) String() string {
	return fmt.Sprintf("%-15s: %s", p.title, p.price)
}

func (p *product) discount(ratio float64) {
	p.price *= money(1 - ratio)
}
