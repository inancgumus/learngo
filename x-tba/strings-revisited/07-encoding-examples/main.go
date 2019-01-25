// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"unicode/utf8"
	"unsafe"
)

func main() {
	// try yourself: try other runes!
	// you can find more here: https://unicode-table.com
	// r := '🙉'

	// r := '\u011e'
	r := 'Ğ'

	// only codepoint (can't be printed)
	fmt.Printf("before encoding: %d\n", r)
	fmt.Printf("  bits : %016b\n", r)
	fmt.Printf("  bytes: % x\n", r)

	// utf-8 encoded string
	encoded := string(r)
	encodedBytes := []byte(encoded)

	fmt.Println()
	fmt.Printf("after encoding: %q\n", encoded)
	fmt.Printf("  bits : %8b\n", encodedBytes)
	fmt.Printf("  bytes: % x\n", encodedBytes)

	// utf-8 string efficient to store and transmit
	//   but, it's harder to use.
	//
	// rune slice is inefficient.
	//   but, it's easy to use.
	fmt.Println()
	fmt.Println("string (utf-8) vs []rune (unicode)")

	s := "hava çok güzel 😳"
	fmt.Printf("%q\n", s)
	fmt.Printf("  size      : %d bytes\n", len(s))
	fmt.Printf("  len       : %d chars\n", utf8.RuneCountInString(s))
	fmt.Printf("  s[5]      : %q\n", s[5])
	fmt.Printf("  s[5:7]    : %q\n", s[5:7])

	runes := []rune(s)
	size := int(unsafe.Sizeof(runes[0])) * len(runes)

	fmt.Printf("\n%q\n", runes)
	fmt.Printf("  size      : %d bytes\n", size)
	fmt.Printf("  len       : %d chars\n", len(runes))
	fmt.Printf("  runes[5]  : %q\n", runes[5])
}
