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
	"strings"
)

func main() {
	books := [...]string{"Kafka's Revenge", "Stay Golden", "Everythingship"}

	upper, lower := books, books

	for i := range books {
		upper[i] = strings.ToUpper(upper[i])
		lower[i] = strings.ToLower(lower[i])
	}

	fmt.Printf("books: %q\n", books)
	fmt.Printf("upper: %q\n", upper)
	fmt.Printf("lower: %q\n", lower)
}
