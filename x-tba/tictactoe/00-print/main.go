// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

/*
~ TICTACTOE GAME IN GO ~
+ This example uses the very basics of the Go language.
+ The goal is learning all the basics.
*/
func main() {
	// VERSION #1: String Concat +
	/*
		fmt.Print("" +
			" TIC~TAC~TOE\n" +
			"\n" +
			"/---+---+---\\\n" +
			"| X | O | X |\n" +
			"+---+---+---+\n" +
			"| X | X |   |\n" +
			"+---+---+---+\n" +
			"| O | O | O |\n" +
			"\\---+---+---/\n")
	*/

	// VERSION #2: String Concat +
	/*
		fmt.Println("" +
			" TIC~TAC~TOE\n" +
			"\n" +
			"/---+---+---\\\n" +
			"| X | O | X |\n" +
			"+---+---+---+\n" +
			"| X | X |   |\n" +
			"+---+---+---+\n" +
			"| O | O | O |\n" +
			"\\---+---+---/")
	*/

	// VERSION #3: Raw Literals (multi line strings)
	fmt.Println(`
 TIC~TAC~TOE

/---+---+---\
| X | O | X |
+---+---+---+
| X | X |   |
+---+---+---+
| O | O | O |
\---+---+---/`)
}
