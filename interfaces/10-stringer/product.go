// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
)

type product struct {
	title    string
	price    money
	released timestamp
}

func (p *product) String() string {
	return fmt.Sprintf("%s: %s (%s)", p.title, p.price, p.released)
}

func (p *product) discount(ratio float64) {
	p.price *= money(1 - ratio)
}
