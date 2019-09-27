// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"strings"
)

type item interface {
	sum() money
	discount(ratio float64)
	fmt.Stringer // same as: `String() string`
}

type list []item

func (l list) String() string {
	if len(l) == 0 {
		return "Sorry. We're waiting for delivery 🚚."
	}

	var str strings.Builder
	for _, it := range l {
		fmt.Fprintf(&str, "%s\n", it)
	}
	fmt.Fprintf(&str, "\tTOTAL  : $%.2f\n", l.sum())

	return str.String()
}

func (l list) sum() (total money) {
	for _, it := range l {
		total += it.sum()
	}
	return
}

func (l list) discount(ratio float64) {
	for _, it := range l {
		it.discount(ratio)
	}
}
