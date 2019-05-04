// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	coders := []string{"pike", "thompson", "lovelace", "knuth"}
	gophers := make(map[string]bool, len(coders))

	gophers["thompson"] = true
	gophers["pike"] = true

	fmt.Println(gophers["pike"])
	fmt.Println(gophers["lovelace"])

	// unnecessary
	// if _, ok := gophers["pike"]; ok {
	// 	fmt.Println("pike is a gopher")
	// } else {
	// 	fmt.Println("pike is NOT a gopher")
	// }

	// easier
	if gophers["pike"] {
		fmt.Println("pike is a gopher")
	} else {
		fmt.Println("pike is NOT a gopher")
	}
}
