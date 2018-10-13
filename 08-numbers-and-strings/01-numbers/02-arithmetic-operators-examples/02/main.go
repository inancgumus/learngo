// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	// what's the value of the ratio?
	// 3 / 2 = 1.5?
	var ratio float64 = 3 / 2
	fmt.Println(ratio)

	// explain
	// above expression equals to this:
	ratio = float64(int(3) / int(2))
	fmt.Println(ratio)

	// how to fix it?
	//
	// remember, when one of the values is a float value
	// the result becomes a float
	ratio = float64(3) / 2
	fmt.Println(ratio)

	// or
	ratio = 3.0 / 2
	fmt.Println(ratio)
}
