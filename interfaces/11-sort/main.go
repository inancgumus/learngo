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
	"sort"
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
	sort.Sort(sort.Reverse(byReleaseDate(l)))

	fmt.Print(l)
}

/*
Summary:

- sort.Sort() can sort any type that implements the sort.Interface

- sort.Interface has three methods: Len(), Less(), Swap()
  - Len() returns the length of a collection.
  - Less(i, j) should return true when an element comes before another one.
  - Swap(i, j)s the elements when Less() returns true.

- sort.Reverse() can reverse sort a type that satisfies the sort.Interface.

- You can customize the sorting:
  - Either by implementing the sort.Interface methods,
  - or by anonymously embedding a type that already satisfies the sort.Interface
  - and adding a Less() method.

- Anonymous embedding means auto-forwarding method calls to an embedded value.
*/
