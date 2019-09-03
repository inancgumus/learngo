// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
)

type book struct {
	product

	// Published (timestamp) knows how to print, encode and decode itself.
	Published timestamp
}

func (b *book) String() string {
	return fmt.Sprintf("%s - (%s)", &b.product, b.Published)
}
