// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// THIS EXERCISE IS OPTIONAL

// ---------------------------------------------------------
// EXERCISE: Print hexes
//
//  1. Print 0 to 9 in hexadecimal
//
//  2. Print 10 to 15 in hexadecimal
//
//  3. Print 17 in hexadecimal
//
//  4. Print 25 in hexadecimal
//
//  5. Print 50 in hexadecimal
//
//  6. Print 100 in hexadecimal
//
// EXPECTED OUTPUT
//  0 1 2 3 4 5 6 7 8 9
//  10 11 12 13 14 15
//  17
//  25
//  50
//  100
//
// NOTES
//  https://stackoverflow.com/questions/910309/how-to-turn-hexadecimal-into-decimal-using-brain
//
// https://simple.wikipedia.org/wiki/Hexadecimal_numeral_system
//
// ---------------------------------------------------------

func main() {
	// EXAMPLES:

	// I'm going to print 10 in hexadecimal
	fmt.Println(0xa)

	// I'm going to print 16 in hexadecimal
	// 0x10
	//   ^^-----  1 * 0 = 0
	//   |
	//   +------ 16 * 1 = 16
	//                  = 16
	fmt.Println(0x10)

	// I'm going to print 150 in hexadecimal
	// 0x96
	//   ^^-----  1 * 6 = 6
	//   |
	//   +------ 16 * 9 = 144
	//                  = 150
	fmt.Println(0x96)

	// COMMENT-OUT ALL THE CODE ABOVE, THEN,
	// ADD YOUR OWN SOLUTIONS BELOW
}
