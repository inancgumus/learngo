package main

// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Please type a search word.")
		return
	}
	query := args[0]

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	words := make(map[string]bool)
	for in.Scan() {
		words[in.Text()] = true
	}

	if words[query] {
		fmt.Printf("The input contains %q.\n", query)
		return
	}
	fmt.Printf("Sorry. The input does contain %q.\n", query)

	// query := "sun"
	// fmt.Println("sun:", words["sun"], "tesla:", words["tesla"])

	// unnecessary
	// if _, ok := words[query]; ok {
	// 	fmt.Printf("The input contains %q.\n", query)
	// 	return
	// }
	// fmt.Printf("Sorry. The input does contain %q.\n", query)
}
