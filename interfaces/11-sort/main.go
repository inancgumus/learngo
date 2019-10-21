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

func main() {
	l := list{
		{title: "moby dick", price: 10, released: toTimestamp(118281600)},
		{title: "odyssey", price: 15, released: toTimestamp("733622400")},
		{title: "hobbit", price: 25},
	}

	// sort.Sort(l)
	// sort.Sort(sort.Reverse(l))
	// sort.Sort(byReleaseDate(l))
	// sort.Sort(sort.Reverse(byReleaseDate(l)))

	fmt.Print(l)
}

/*
Summary:

- sort.Sort() can sort any type that implements the sort.Interface

- sort.Interface has three methods: Len(), Less(), Swap()
  - Len() returns the length of a collection.
  - Less() should return true when an element comes before another one.
  - Swap()s the elements when Less() returns true.

- sort.Reverse() sorts a sort.Interface value.

- You can customize the sorting:
  - by anonymously embedding the sort.Interface type
  - and adding a Less() method.

- Anonymous embedding means auto-forwarding method calls to an embedded value.
*/
