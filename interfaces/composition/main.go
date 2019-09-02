// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"encoding/json"
	"fmt"
)

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

	out, _ := json.MarshalIndent(store, "", "\t")
	fmt.Println(string(out))

	// store.discount(.5)
	// fmt.Print(store)

	// var (
	// 	nilItem item
	// 	nilBook *book
	// )

	// fmt.Println("nilBook     ?", nilBook == nil)
	// fmt.Println("nilItem?", nilItem == nil)

	// nilItem = nilBook
	// fmt.Println("nilItem?", nilItem == nil)

	// books := store[:3]

	// books.discount(.5)
	// fmt.Printf("%s\n", store)

	// others := store[3:]
	// sort.Sort(byPrice(others))
	// fmt.Printf("\n%s\n", store)

	// store = list{books, others}
	// fmt.Printf("\n%s\n", store)

	// // sort.Sort(sort.Reverse(byPrice(store)))
	// sort.Sort(byName(store))
}
