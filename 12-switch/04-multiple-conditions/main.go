// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"os"
)

func main() {
	city := os.Args[1]

	switch city {
	case "Paris", "Lyon":
		fmt.Println("France")
	case "Tokyo":
		fmt.Println("Japan")
	}

	// if city == "Paris" || city == "Lyon" {
	// 	fmt.Println("France")
	// } else if city == "Tokyo" {
	// 	fmt.Println("Japan")
	// }
}
