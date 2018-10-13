// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	score, valid := 3, true

	if score > 3 && valid {
		fmt.Println("good")
	} else {
		fmt.Println("low")
	}
}
