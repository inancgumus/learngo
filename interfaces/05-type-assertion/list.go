// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

type printer interface {
	print()

	// use type assertion when you cannot change the interface.
	// discount(ratio float64)
}

type list []printer

func (l list) print() {
	if len(l) == 0 {
		fmt.Println("Sorry. We're waiting for delivery 🚚.")
		return
	}

	for _, it := range l {
		it.print()
	}
}

// type assertion can extract the wrapped value.
// or: it can put the value into another interface.
func (l list) discount(ratio float64) {
	// you can declare an interface in a function/method as well.
	// interface is just a type.
	type discounter interface {
		discount(float64)
	}

	for _, it := range l {
		// you can assert to an interface.
		// and extract another interface.
		if it, ok := it.(discounter); ok {
			it.discount(ratio)
		}
	}
}

// ----

// func (l list) discount(ratio float64) {
// 	for _, it := range l {
// 		// you can inline-assert interfaces
// 		// interface is just a type.
// 		g, ok := it.(interface{ discount(float64) })
// 		if !ok {
// 			continue
// 		}

// 		g.discount(ratio)
// 	}
// }

// ----

// // type assertion can pull out the real value behind an interface value.
// // it can also check whether the value convertable to a given type.
// func (l list) discount(ratio float64) {
// 	for _, it := range l {
// 		g, ok := it.(*game)
// 		if !ok {
// 			continue
// 		}

// 		g.discount(ratio)
// 	}
// }
