// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

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
	// sort.Sort(byRelease{l})
	sort.Sort(sort.Reverse(byRelease{l}))

	fmt.Print(l)
}
