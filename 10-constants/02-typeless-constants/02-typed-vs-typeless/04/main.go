// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	const min = 42

	// i've removed int from the below declaration
	// since, min's default type is int (you'll learn)
	var i = min

	var f float64 = min
	var b byte = min
	var j int32 = min
	var r rune = min

	// behind the scenes:
	// below statement equals to:
	//
	// var b byte = min
	b = byte(min)

	fmt.Println(i, f, b, j, r)
}
