// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("[english word] -> [turkish word]")
		return
	}
	query := args[0]

	dict := map[string]string{
		"good":    "iyi",
		"great":   "harika",
		"perfect": "mükemmel",

		// #5: this overwrites the mükemmel in the turkish map
		"awesome": "mükemmel",
	}

	// #1: map value is a pointer to a real map in the memory
	// turkish := dict
	// turkish["good"] = "güzel"
	// dict["great"] = "kusursuz"

	// #6: delete can delete a pair from a map
	delete(dict, "awesome")

	// #7: no-op
	delete(dict, "awesome")
	delete(dict, "notexisting")

	// #8: destroying a map completely (false)
	// dict = nil

	// #9: destroying a map completely (true)
	//     replaces the whole loop with a single instruction:
	//     "call runtime.mapclear()"
	for k := range dict {
		delete(dict, k)
	}

	// #2: make initializes a new map value and returns a pointer to it
	// turkish := make(map[string]string)

	// #3: make(T, hint) => hint is an advice for initing a large enough map
	turkish := make(map[string]string, len(dict))
	for k, v := range dict {
		turkish[v] = k
	}

	// #4: add turkish dictionary
	// fmt.Printf("english: %q\nturkish: %q\n", dict, turkish)

	if value, ok := dict[query]; ok {
		fmt.Printf("%q means %#v\n", query, value)
		return
	}

	if value, ok := turkish[query]; ok {
		fmt.Printf("%q means %#v\n", query, value)
		return
	}

	fmt.Printf("%q not found.\n", query)
}
