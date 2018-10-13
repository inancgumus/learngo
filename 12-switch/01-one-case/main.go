// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

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
