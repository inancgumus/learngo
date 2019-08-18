// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
)

func main() {
	nums := []int{56, 89, 15, 25, 30, 50}

	// ----------------------------------------
	// breaks the connection:
	// mine and nums now have different backing arrays

	// verbose solution:
	// var mine []int
	// mine = append(mine, nums[:3]...)

	// better solution (almost the same thing):
	mine := append([]int(nil), nums[:3]...)
	// ----------------------------------------

	mine[0], mine[1], mine[2] = -50, -100, -150
	fmt.Println("Mine         :", mine)
	fmt.Println("Original nums:", nums[:3])
}
