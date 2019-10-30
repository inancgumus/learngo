// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

func main() {
	prev := [3]string{"Kafka's Revenge", "Stay Golden", "Everythingship"}

	// You can't do this:
	// books = prev

	var books [4]string

	for i, b := range prev {
		books[i] += b + " 2nd Ed."
	}

	books[3] = "Awesomeness"

	fmt.Printf("last year:\n%#v\n", prev)
	fmt.Printf("\nthis year:\n%#v\n", books)
}
