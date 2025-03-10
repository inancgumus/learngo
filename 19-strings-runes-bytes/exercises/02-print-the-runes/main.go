// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Print the runes
//
//  1. Loop over the "console" word and print its runes one by one,
//     in decimals, hexadecimals and binary.
//
//  2. Manually put the runes of the "console" word to a byte slice, one by one.
//
//     As the elements of the byte slice use only the rune literals.
//
//     Print the byte slice.
//
//  3. Repeat the step 2 but this time, as the elements of the byte slice,
//     use only decimal numbers.
//
//  4. Repeat the step 2 but this time, as the elements of the byte slice,
//     use only hexadecimal numbers.
//
//
// EXPECTED OUTPUT
//   Run the solution to see the expected output.
// ---------------------------------------------------------

func main() {
	const word = "console"

	for _, c := range word {
		fmt.Printf("%c\n", c)
		fmt.Printf("\tdecimal: %d\n", c)
		fmt.Printf("\thex    : 0x%x\n", c)
		fmt.Printf("\tbinary : 0b%b\n", c)
	}

	fmt.Printf("with runes       : %s\n",
		[]byte{'c', 'o', 'n', 's', 'o', 'l', 'e'})
	fmt.Printf("with decimals    : %s\n",
		[]byte{99, 111, 110, 115, 111, 108, 101})
	fmt.Printf("with hex         : %s\n",
		[]byte{0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65})

}
