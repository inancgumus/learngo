// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	var (
		width  uint8 = 255
		height       = 255 // int
	)

	width++

	if int(width) < height {
		fmt.Println("height is greater")
	}

	fmt.Printf("width: %d height: %d\n", width, height)
}
