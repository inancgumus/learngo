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
	Title    string    `json:"title"`
	Price    money     `json:"price"`
	Released timestamp `json:"released"`
}

func (p *product) String() string {
	return fmt.Sprintf("%s: %s (%s)", p.Title, p.Price, p.Released)
}

func (p *product) discount(ratio float64) {
	p.Price *= money(1 - ratio)
}
