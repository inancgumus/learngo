// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	var g, o rune

	g, o = 'g', 'o'
	g, o = 103, 111
	g, o = 0x67, 0x6f
	g, o = '\U00000067', '\U0000006f'
	g, o = '\u0067', '\u006f'
	g, o = '\x67', '\x6f'

	fmt.Println("codepoints")
	fmt.Printf("  dec       : %d %d\n", g, o)
	fmt.Printf("  hex       : %x %x\n", g, o)
	fmt.Printf("  unicode   : %U %U\n", g, o)
	fmt.Printf("  chars     : %c %c\n", g, o)

	// g++
	// o -= 6

	g -= 'a' - 'A'
	o -= 'a' - 'A'

	fmt.Println("codepoints")
	fmt.Printf("  dec       : %d %d\n", g, o)
	fmt.Printf("  hex       : %x %x\n", g, o)
	fmt.Printf("  unicode   : %U %U\n", g, o)
	fmt.Printf("  chars     : %c %c\n", g, o)

	// string representations
	// fmt.Print("string()    : ", string(g), string(o), "\n")
	// fmt.Print("hex 1 byte  : \x67\x6f \n")
	// fmt.Print("hex 2 bytes : \u0067\u006f \n")
	// fmt.Print("hex 4 bytes : \U00000067\U0000006f \n")
}
