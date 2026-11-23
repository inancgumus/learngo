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
	"unicode/utf8"
)

// ---------------------------------------------------------
// EXERCISE: Rune Manipulator
//
//  Please read the comments inside the following code.
//
// EXPECTED OUTPUT
//  Please run the solution.
// ---------------------------------------------------------

func main() {
	words := []string{
		"cool",
		"güzel",
		"jīntiān",
		"今天",
		"read 🤓",
	}

	for _, word := range words {
		fmt.Printf("%q\n", word)
		fmt.Printf("\t has %d bytes and %d runes\n",
			len(word), utf8.RuneCountInString(word))
		fmt.Printf("\t bytes   :% x \n",
			word)
		fmt.Print("\t runes   :")

		for _, r := range word {
			fmt.Printf(" %x", r)
		}
		fmt.Println()

		fmt.Print("\trunes   :")
		for _, r := range word {
			fmt.Printf("%q", r)
		}
		fmt.Println()

		r1, size := utf8.DecodeRuneInString(word)

		fmt.Printf("\tfirst   : %q (%d bytes)\n", r1, size)

		r2, size := utf8.DecodeLastRuneInString(word)

		fmt.Printf("\tlast   : %q (%d bytes)\n", r2, size)

		_, first := utf8.DecodeRuneInString(word)
		_, second := utf8.DecodeRuneInString(word[first:])
		fmt.Printf("\tfirst 2 : %q\n", word[:first+second])

		_, last1 := utf8.DecodeLastRuneInString(word)
		_, last2 := utf8.DecodeLastRuneInString(word[:len(word)-last1])
		fmt.Printf("\tlast 2  : %q\n", word[len(word)-last2-last1:])

		r_slice := []rune(word)

		fmt.Printf("\t first2   :%q \n", string(r_slice[:2]))
		fmt.Printf("\t last 2   :%q \n", string(r_slice[len(r_slice)-2:]))

	}
}
