// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	books := [4]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
		"Kafka's Revenge 2nd Edition",
	}

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Tell me a book title")
		return
	}
	query := strings.ToLower(args[0])

	fmt.Println("Search Results:")

	var found bool
	for _, v := range books {
		if strings.Contains(strings.ToLower(v), query) {
			fmt.Println("+", v)
			found = true
		}
	}

	if !found {
		fmt.Printf("We don't have the book: %q\n", query)
	}
}
