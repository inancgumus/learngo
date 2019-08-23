// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

type printer interface {
	print()
	// discount(ratio float64)
}

type list []printer

func (l list) print() {
	if len(l) == 0 {
		fmt.Println("Sorry. Our store is closed. We're waiting for the delivery 🚚.")
		return
	}

	for _, it := range l {
		it.print()
	}
}

//
// type switch can find the type behind an interface value.
//
func (l list) discount(ratio float64) {
	type discounter interface {
		discount(float64)
	}

	for _, it := range l {
		if it, ok := it.(discounter); ok {
			it.discount(ratio)
		}
	}
}

// ----

//
// type assertion can pull out the real value behind an interface value.
//
// it can also check whether the value convertable to a given type.
//
// func (l list) discount(ratio float64) {
// 	for _, it := range l {
// 		g, ok := it.(interface{ discount(float64) })
// 		if !ok {
// 			continue
// 		}

// 		g.discount(ratio)
// 	}
// }

// ----

//
// type switch can pull out the real value behind an interface value.
//
// func (l list) discount(ratio float64) {
// 	for _, it := range l {
// 		switch it := it.(type) {
// 		case *game:
// 			it.discount(ratio)
// 		case *puzzle:
// 			it.discount(ratio)
// 		}
// 	}
// }

// ----

//
// type switch can find the type behind an interface value.
//
// func (l list) discount(ratio float64) {
// 	for _, it := range l {
// 		switch it.(type) {
// 		case *game:
// 			fmt.Print("it's a *game!   --> ")
// 		case *puzzle:
// 			fmt.Print("it's a *puzzle! --> ")
// 		default:
// 			fmt.Print("neither         --> ")
// 		}
// 		it.print()
// 	}
// }
