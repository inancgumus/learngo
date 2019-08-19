// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

// TODO: NEW
type printer interface {
	print()
}

// TODO: NEW
type pricer interface {
	total() float64
}

// TODO: NEW
// interface embedding
type item interface {
	printer
	pricer
}

// a slice of item interface values
type list []item

func (l list) list() {
	// `list` acts like a `[]game`
	if len(l) == 0 {
		fmt.Println("Sorry. Our store is closed. We're waiting for the delivery 🚚.")
		return
	}

	fmt.Printf("My Store:\n")
	fmt.Printf("---------\n")
	for _, it := range l {
		it.print()
	}

	// TODO: NEW
	fmt.Printf("\nGROSS: $%.2f\n", l.total())
}

// TODO: NEW
func (l list) total() (gross float64) {
	for _, it := range l {
		gross += it.total()
	}
	return gross
}

// TODO: NEW
func (l list) discount(ratio float64) {
	type discounter interface {
		discount(float64)
	}

	for _, it := range l {
		dit, ok := it.(discounter)
		if !ok {
			continue
		}

		dit.discount(ratio)
	}
}
