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
	"unicode"
)

func main() {
	s := "Func Values Are Awesome..."

	// s = strings.Map(unicode.ToUpper, s)
	// s = strings.Map(unicode.ToLower, s)
	// s = strings.Map(excite, s)
	// s = strings.Map(calm, s)

	// fmt.Printf("%T\n", strings.Map)
	s = mapx(calm, s)

	fmt.Println(s)
}

func mapx(mapping func(rune) rune, s string) string {
	var ns []rune

	for _, r := range s {
		if r = mapping(r); r == -1 {
			continue
		}
		ns = append(ns, r)
	}
	return string(ns)
}

func calm(r rune) rune {
	if unicode.IsPunct(r) {
		return -1
	}
	return unicode.ToLower(r)
}

func excite(r rune) rune {
	if r == '.' {
		return '!'
	}
	return unicode.ToTitle(r)
}
