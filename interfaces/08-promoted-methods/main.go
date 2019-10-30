// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

func main() {
	store := list{
		&book{product{"moby dick", 10}, 118281600},
		&book{product{"odyssey", 15}, "733622400"},
		&book{product{"hobbit", 25}, nil},
		&puzzle{product{"rubik's cube", 5}},
		&game{product{"minecraft", 20}},
		&game{product{"tetris", 5}},
		&toy{product{"yoda", 150}},
	}

	store.discount(.5)
	store.print()

	// t := &toy{product{"yoda", 150}}
	// fmt.Printf("%#v\n", t)

	// b := &book{product{"moby dick", 10}, 118281600}
	// fmt.Printf("%#v\n", b)
}
