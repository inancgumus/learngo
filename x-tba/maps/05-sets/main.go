// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	// set is a data structure that contain the same element only once.

	// maps can also be used as sets.

	// let's create a slice of coders.
	coders := []string{"pike", "thompson", "lovelace", "knuth"}

	// let's create a map for tracking whether they are gopher or not.
	gophers := make(map[string]bool, len(coders))

	// let's mark the gophers
	gophers["thompson"] = true
	gophers["pike"] = true

	// print the gophers
	for _, coder := range coders {
		// i see this a lot, especially in beginner code.
		//
		// you don't have to use the "comma, ok" idiom.
		// m[k] = false for non-existing bool elements.
		// use it to your advantage.
		//
		// DONT:
		// if _, ok := gophers[coder]; ok { ... }
		//
		// DO:
		// for this trick to work, you should set all the map elements to true
		if gophers[coder] {
			fmt.Printf("%s is a gopher.\n", coder)

			// don't use else. stay at the same indentation level.
			continue
		}

		fmt.Printf("%s is not a gopher.\n", coder)
	}
}
