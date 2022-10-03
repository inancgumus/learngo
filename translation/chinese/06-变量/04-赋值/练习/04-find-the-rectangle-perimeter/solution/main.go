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
		perimeter     int
		width, height = 5, 6
	)

	// 首先计算: (宽 + 高)
	// 然后            :  用 2 与它相乘

	// 就和在数学里一样

	perimeter = 2 * (width + height)

	fmt.Println(perimeter)
}
