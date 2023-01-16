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
	// 以下方案是正确的:
	x := 5. / 2
	// x := 5 / 2.
	// x := float64(5) / 2
	// x := 5 / float64(2)

	fmt.Println(x)
}
