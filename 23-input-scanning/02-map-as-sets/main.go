// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Please type a search word.")
		return
	}
	query := args[0]

	rx := regexp.MustCompile(`[^a-z]+`)

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	words := make(map[string]bool)
	for in.Scan() {
		word := strings.ToLower(in.Text())
		word = rx.ReplaceAllString(word, "")

		if len(word) > 2 {
			words[word] = true
		}
	}

	// for word := range words {
	// 	fmt.Println(word)
	// }

	if words[query] {
		fmt.Printf("The input contains %q.\n", query)
		return
	}
	fmt.Printf("Sorry. The input does not contain %q.\n", query)

	// query := "sun"
	// fmt.Println("sun:", words["sun"], "tesla:", words["tesla"])

	// unnecessary
	// if _, ok := words[query]; ok {
	// 	fmt.Printf("The input contains %q.\n", query)
	// 	return
	// }
	// fmt.Printf("Sorry. The input does not contain %q.\n", query)
}
