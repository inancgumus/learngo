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
	"strconv"
	"time"
)

type book struct {
	// embed the product type into the book type.
	// all the fields and methods of the product will be
	// available in this book type.
	product
	published interface{}
}

// book type's print method takes priority.
//
// + when you call it on a book value, the following method will
//   be executed.
//
// + if it wasn't here, the product type's print method would
//   have been executed.
func (b *book) print() {
	// the book can also call the embedded product's print method
	// if it wants to, as in here:
	b.product.print()

	p := format(b.published)
	fmt.Printf("\t - (%v)\n", p)
}

func format(v interface{}) string {
	var t int

	switch v := v.(type) {
	case int:
		// book{title: "moby dick", price: 10, published: 118281600},
		t = v
	case string:
		// book{title: "odyssey", price: 15, published: "733622400"},
		t, _ = strconv.Atoi(v)
	default:
		// book{title: "hobbit", price: 25},
		return "unknown"
	}

	// Mon Jan 2 15:04:05 -0700 MST 2006
	const layout = "2006/01"

	u := time.Unix(int64(t), 0)
	return u.Format(layout)
}
