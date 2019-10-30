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
	var (
		width  uint16
		height uint16
	)

	width, height = 255, 265
	width += 10

	fmt.Printf("width: %d height: %d\n", width, height)
	fmt.Println("are they equal?", width == height)
}
