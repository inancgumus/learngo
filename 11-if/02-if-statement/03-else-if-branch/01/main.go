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
	score := 3

	if score > 3 {
		fmt.Println("good")
	} else if score == 3 {
		fmt.Println("on the edge")
	} else {
		fmt.Println("low")
	}
}
