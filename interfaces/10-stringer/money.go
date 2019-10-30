// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

type money float64

// String() returns a string representation of money.
// money is an fmt.Stringer.
func (m money) String() string {
	return fmt.Sprintf("$%.2f", m)
}
