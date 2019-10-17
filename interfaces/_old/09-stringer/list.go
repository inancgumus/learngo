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
	discount(ratio float64)

	// ~~ interface embedding ~~
	//
	// item interface embeds the fmt.Stringer interface.
	//
	// Go adds the methods of the fmt.Stringer
	// to this item interface.
	//
	// same as this:
	// String() string
	// fmt.Stringer
}
type list []item

// String method makes the list an fmt.Stringer.
func (l list) String() string {
	if len(l) == 0 {
		return "Sorry. We're waiting for delivery 🚚."
	}

	// use the strings.Builder when you're combining
	// a long list of strings together.
	var str strings.Builder

	for _, it := range l {
		// the builder.WriteString doesn't know about the stringer interface.
		// because it takes a string argument.
		// so, you need to call the String method yourself.
		// str.WriteString(it.String())
		// str.WriteRune('\n')

		// or slower way:
		// Print funcs can detect fmt.Stringer automatically.
		fmt.Fprintln(&str, it)
	}

	return str.String()
}

func (l list) discount(ratio float64) {
	for _, it := range l {
		it.discount(ratio)
	}
}
