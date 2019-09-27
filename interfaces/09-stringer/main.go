// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	// The money type is a stringer.
	// You don't need to call the String method when printing a value of it.
	// var pocket money = 10
	// fmt.Println(pocket)

	store := list{
		&book{product{"moby dick", 10}, toTimestamp(118281600)},
		&book{product{"odyssey", 15}, toTimestamp("733622400")},
		&book{product{"hobbit", 25}, toTimestamp(nil)},
		&puzzle{product{"rubik's cube", 5}},
		&game{product{"minecraft", 20}},
		&game{product{"tetris", 5}},
		&toy{product{"yoda", 150}},
	}

	store.discount(.5)

	// The list is a stringer.
	// The `fmt.Print` function can print the `store`
	// by calling `store`'s `String()` method.
	//
	// Underneath, `fmt.Print` uses a type switch to
	// detect whether a type is a Stringer:
	// https://golang.org/src/fmt/print.go#L627
	fmt.Print(store)
}
