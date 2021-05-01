// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func main() {
	word := []byte("öykü")
	fmt.Printf("%s = % [1]x\n", word)

	// how to make the first rune uppercase?

	// you need to find the starting and ending position of the first rune

	// 1st way: `for range`
	// you can't get the runes by ranging over a byte slice
	// first, you need to convert it to a string
	var size int
	for i := range string(word) {
		if i > 0 {
			size = i
			break
		}
	}

	// 2nd way: let's do it using the utf8 package's DecodeRune function
	_, size = utf8.DecodeRune(word)

	// overwrite the current bytes with the new uppercased bytes
	copy(word[:size], bytes.ToUpper(word[:size]))

	// to get printed bytes/runes need to be encoded in a string
	fmt.Printf("%s = % [1]x\n", word)
}
