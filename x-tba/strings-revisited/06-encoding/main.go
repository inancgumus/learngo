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

// Please run this code and experiment with it
// Observe the results

// USELESS-NOTE : "Öykü" means "Story" in Turkish!

func main() {
	fmt.Println("-----------------------------------")
	fmt.Println("ASCII Codepoints")
	fmt.Println("-----------------------------------")

	var (
		a, z   byte = 'a', 'z'
		A, Z   byte = 'A', 'Z'
		d0, d9 byte = '0', '9'
	)

	for _, c := range []byte{a, z, A, Z, d0, d9} {
		fmt.Printf("%c - 1 byte - %[1]U - %[1]d\n", c)
	}

	fmt.Println("\n-----------------------------------")
	fmt.Println("Unicode Codepoints")
	fmt.Println("-----------------------------------")

	var (
		Ö      = 'Ö'
		栗      = '栗'
		monkey = '🙉'
	)
	for _, c := range []rune{rune(A), Ö, 栗, monkey} {
		fmt.Printf("%c - %d bytes - %[1]U - %[1]d\n", c, cptb(c))
	}

	fmt.Println("\n-----------------------------------")
	fmt.Println("UTF-8 Encoded")
	fmt.Println("-----------------------------------")

	// utf8.RuneLen finds the number of bytes necessary for
	// encoding a codepoint to utf8
	for _, c := range []rune{rune(A), Ö, 栗, monkey} {
		fmt.Printf("%c - %d bytes - %[1]U - %[1]d\n", c,
			utf8.RuneLen(c))
	}

	fmt.Println("\n-----------------------------------")
	fmt.Println("Example: Unicode Codepoints")
	fmt.Println("-----------------------------------")

	var (
		ö = 'ö'
		y = 'y'
		k = 'k'
		ü = 'ü'
	)

	var (
		oykuRunes = []rune{ö, y, k, ü}
		total     int
	)

	for _, c := range oykuRunes {
		fmt.Printf("%c - %d bytes - %[1]U - %[1]d\n", c, cptb(c))

		// unsafe.Sizeof finds the memory size of simple values
		// don't use it in production-level code -> it's unsafe!
		total += int(unsafe.Sizeof(c))
	}
	fmt.Printf("TOTAL: %d bytes.\n", total)

	fmt.Println("\n-----------------------------------")
	fmt.Println("Example: Indexing")
	fmt.Println("-----------------------------------")

	fmt.Printf("%c%c%c%c\n",
		oykuRunes[0], oykuRunes[1], oykuRunes[2],
		oykuRunes[len(oykuRunes)-1])

	// string to []rune
	oykuRunes = []rune("öykü")
	fmt.Printf("%c%c%c%c\n",
		oykuRunes[0], oykuRunes[1], oykuRunes[2],
		oykuRunes[len(oykuRunes)-1])

	fmt.Println("\n-----------------------------------")
	fmt.Println("Example: UTF-8 Encoding")
	fmt.Println("-----------------------------------")

	// this is also ok
	// oykuString := string(oykuRunes)

	oykuString := "öykü"

	fmt.Printf("TOTAL bytes in oykuRunes : %d\n", total)
	fmt.Printf("TOTAL bytes in oykuString: %d\n", len(oykuString))
	fmt.Printf("TOTAL runes in oykuString: %d\n",
		utf8.RuneCountInString(oykuString))

	fmt.Printf("Runes of oykuString      : %s\n", oykuString)
	fmt.Printf("Bytes of oykuString      : % x\n", oykuString)

	fmt.Println()
	for i := 0; i < len(oykuString); i++ {
		fmt.Printf("oykuString[%d]: %c\n", i, oykuString[i])
	}

	// slicing returns a slice with the type of the sliced value
	// so, the sliced value is a string, then a string is returned
	//
	// example:
	// oykuString[0:2] is a string
	fmt.Println()
	fmt.Printf("oykuString[0:2]: %q\n", oykuString[0:2])
	fmt.Printf("oykuString[4:6]: %q\n", oykuString[4:6])
}

// -------------------------------------------------------------------
// cptb finds how many bytes are necessary to represent a codepoint
// cptb means codepoint to bytes
func cptb(r rune) int {
	switch {
	case r <= 0xFF: // 255
		return 1
	case r <= 0xFFFF: // 65,535
		return 2
	case r <= 0xFFFFF: // 16,777,215
		return 3
	}
	return 4
}
