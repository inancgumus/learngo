// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

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
