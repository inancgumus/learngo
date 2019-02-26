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
	/*
		#1: Nil Map: Read-Only
	*/
	var dict map[string]string
	fmt.Printf("Zero-value of a map: %#v\n", dict)

	/*
		#5: You cannot assign to a nil map.
	*/
	// dict["up"] = "yukarı"
	// dict["down"] = "aşağı"

	/*
		#2: Map retrieval is O(1) — on average.
	*/
	key := "good"
	//
	// #4: you can use an uninitialized map without checking it is nil
	//
	// if dict != nil {
	value := dict[key]
	fmt.Printf("%q means %#v\n", key, value)
	// }

	/*
		#3: Cannot use non-comparable types as map key types
	*/
	// var broken map[[]int]int
	// var broken map[map[int]string]bool
	//
	// A map can only be compared to nil value
	// _ = dict == nil

	/*
		#1 Step 2
	*/
	fmt.Printf("The dictionary contains %d words.\n", len(dict))
}
