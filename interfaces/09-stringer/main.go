// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	store := list{
		&book{product{"moby dick", 10}, toTimestamp(118281600)},
		&book{product{"odyssey", 15}, toTimestamp("733622400")},
		&book{product{"hobbit", 25}, unknown},
		&puzzle{product{"rubik's cube", 5}},
		&game{product{"minecraft", 20}},
		&game{product{"tetris", 5}},
		&toy{product{"yoda", 150}},
	}

	store.discount(.5)

	// the store doesn't have a print method anymore.
	// but the Print function can print it.
	// it's because, the list satisfies the stringer.
	fmt.Print(store)

	// timestamp is useful even if it's zero.
	var ts timestamp
	fmt.Println(ts)
}
