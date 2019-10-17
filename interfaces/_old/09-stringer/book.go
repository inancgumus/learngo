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
	published timestamp
}

// String method makes the book an fmt.Stringer.
// The list was calling the embedded product type's String(),
// now it calls the book.String().
func (b *book) String() string {
	// product.String() has a pointer receiver.
	// Therefore, you need to manually take the product's address.
	//
	// If you pass: "b.product", Go would pass it as a copy to `Sprintf`.
	// In that case, Go can't deference b.product automatically.
	// It's because: b.product would be a different value, a copy.
	return fmt.Sprintf("%s - (%s)", &b.product, b.published)
}
