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
	"unsafe"
)

func main() {
	str := "Yūgen ☯ 💀"

	// can't change a string
	// a string is a read-only byte-slice
	// str[0] = 'N'
	// str[1] = 'o'

	bytes := []byte(str)

	// can change a byte slice
	// bytes[0] = 'N'
	// bytes[1] = 'o'

	str = string(bytes)

	fmt.Printf("%s\n", str)
	fmt.Printf("\t%d bytes\n", len(str))
	fmt.Printf("\t%d runes\n", utf8.RuneCountInString(str))
	fmt.Printf("% x\n", bytes)
	fmt.Printf("\t%d bytes\n", len(bytes))
	fmt.Printf("\t%d runes\n", utf8.RuneCount(bytes))

	// fmt.Println()
	// for i, r := range str {
	// 	fmt.Printf("str[%2d] = % -12x = %q\n", i, string(r), r)
	// }

	fmt.Println()
	fmt.Printf("1st byte   : %c\n", str[0])           // ok
	fmt.Printf("2nd byte   : %c\n", str[1])           // not ok
	fmt.Printf("2nd rune   : %s\n", str[1:3])         // ok
	fmt.Printf("last rune  : %s\n", str[len(str)-4:]) // ok

	// disadvantage: each one is 4 bytes
	runes := []rune(str)

	fmt.Println()
	fmt.Printf("%s\n", str)
	fmt.Printf("\t%d bytes\n", int(unsafe.Sizeof(runes[0]))*len(runes))
	fmt.Printf("\t%d runes\n", len(runes))

	fmt.Printf("1st rune   : %c\n", runes[0])
	fmt.Printf("2nd rune   : %c\n", runes[1])
	fmt.Printf("first five : %c\n", runes[:5])

	fmt.Println()

	word := "öykü"
	fmt.Printf("%q in runes: %c\n", word, []rune(word))
	fmt.Printf("%q in bytes: % [1]x\n", word)

	fmt.Printf("%s %s\n", word[:2], []byte{word[0], word[1]}) // ö
	fmt.Printf("%c\n", word[2])                               // y
	fmt.Printf("%c\n", word[3])                               // k
	fmt.Printf("%s %s\n", word[4:], []byte{word[4], word[5]}) // ü
}
