// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

type list []item

func (l list) print() {
	if len(l) == 0 {
		fmt.Println("Sorry. Our store is closed. We're waiting for the delivery 🚚.")
		return
	}

	for _, it := range l {
		it.print()
	}

	// TODO: NEW
	fmt.Printf("\tTOTAL  : $%.2f\n", l.sum())
}

// TODO: NEW
func (l list) sum() (n money) {
	for _, it := range l {
		n += it.sum()
	}
	return n
}
