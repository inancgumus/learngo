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
	/*
		#2—A: Get the key from CLI
	*/
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("[english word] -> [turkish word]")
		return
	}
	query := args[0]

	/*
		#1: Empty Map Literal
	*/
	// dict := map[string]string{}

	/*
		#3: Map Literal
		Creates, initializes and returns a new map
		with the given key-value pairs
	*/
	dict := map[string]string{
		"good":    "kötü",
		"great":   "harika",
		"perfect": "mükemmel",
		/*
			#4: keys and values should be the same type
				(key and value types do not need to be the same)
		*/
		// 42: "forty two",
		// "forty two": 42,
	}

	// you can only add new pairs to an initialized map
	dict["up"] = "yukarı"  // adds a new pair
	dict["down"] = "aşağı" // adds a new pair
	dict["good"] = "iyi"   // #5: overwrites the value at the key: "good"
	dict["mistake"] = ""   // #8: a key with a zero-value

	// you can query an uninitialized map for its length
	fmt.Printf("The dictionary contains %d words.\n", len(dict))

	/*
		#6: getting values from a map using keys directly
	*/
	// fmt.Println("good      -> ", dict["good"])
	// fmt.Println("great     -> ", dict["great"])
	// fmt.Println("perfect   -> ", dict["perfect"])

	/*
		#11: looping over a map might be a sign of a design mistake
			 in your software.

			 (it prints unordered output)
	*/
	// for k, v := range dict {
	// 	fmt.Printf("%q means %#v\n", k, v)
	// }

	/*
		#12: printing a map (ordered output since Go 1.12)
	*/
	// fmt.Printf("%#v\n", dict)

	/*
		#13: compare a map using its printed output
			 (handy when testing)
	*/
	// copied := map[string]string{"up": "yukarı", "down": "aşağı",
	// 	"mistake": "", "good": "iyi", "great": "harika",
	// 	"perfect": "mükemmel"}

	// first := fmt.Sprintf("%s", dict)
	// second := fmt.Sprintf("%s", copied)

	// if first == second {
	// 	fmt.Println("Maps are equal")
	// }

	/*
		#2—B: retrieve values by key - O(1) efficiency
			(even from an uninitialized map)
	*/
	// value := dict[query]

	/*
		#7: check for non-existing key using zero-value
	*/
	// if value == "" {}

	/*
		#9: check for non-existing key: with comma, ok
	*/
	// value, ok := dict[query]
	// if !ok {}

	/*
		#10: comma ok in a short if
	*/
	if value, ok := dict[query]; ok {
		fmt.Printf("%q means %#v\n", query, value)
		return
	}
	fmt.Printf("%q not found.\n", query)
}
