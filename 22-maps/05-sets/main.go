// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	// let's create a slice of coders.
	coders := []string{"pike", "thompson", "lovelace", "knuth"}

	// let's create a map for tracking whether they are gophers or not.
	gophers := make(map[string]bool, len(coders))

	// let's mark the gophers
	gophers["thompson"] = true
	gophers["pike"] = true

	fmt.Println(gophers["pike"])
	fmt.Println(gophers["lovelace"])

	if _, ok := gophers["pike"]; ok {
		fmt.Println("pike is a gopher")
	} else {
		fmt.Println("pike is NOT a gopher")
	}
}
