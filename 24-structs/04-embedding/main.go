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
)

func main() {
	// #1: declare the types
	type text struct {
		title string
		words int
	}

	type book struct {
		// title string
		// words int

		// #3: include the text as a field
		// text text

		// #4: embed the text
		text
		isbn string

		// #5: add a conflicting field
		title string
	}

	// #2: print a book
	// moby := book{title: "moby dick", words: 206052, isbn: "102030"}
	// fmt.Printf("%s has %d words (isbn: %s)\n", moby.title, moby.words, moby.isbn)

	// #3b: type the text in its own field
	moby := book{
		// #5c: type the field in a new field
		// title: "conflict",
		text: text{title: "moby dick", words: 206052},
		isbn: "102030",
	}

	moby.text.words = 1000
	moby.words++

	// // #4b: print the book
	fmt.Printf("%s has %d words (isbn: %s)\n",
		moby.title, // equals to: moby.text.title
		moby.words, // equals to: moby.text.words
		moby.isbn)

	// #3c: print the book
	// fmt.Printf("%s has %d words (isbn: %s)\n",
	// 	moby.text.title, moby.text.words, moby.isbn)

	// #5b: print the conflict
	fmt.Printf("%#v\n", moby)

	// go get -u github.com/davecgh/go-spew/spew
	// spew.Dump(moby)
}
