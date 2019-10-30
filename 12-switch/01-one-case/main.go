// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

func main() {
	city := "Paris"

	switch city {
	case "Paris":
		fmt.Println("France")
	}

	// SIMILAR TO IF
	// ------------------------------------

	// switch statement is converted to an if statement
	// automatically behind the scenes
	//
	// above switch statement is similar to this if

	// if city == "Paris" {
	// 	fmt.Println("France")
	// }
}
