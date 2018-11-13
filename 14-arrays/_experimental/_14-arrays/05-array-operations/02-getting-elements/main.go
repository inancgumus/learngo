// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	ages := [2]int{15, 20}

	// WRONG (try it):
	// fmt.Println(ages[-1])

	fmt.Println("1st item:", ages[0])
	fmt.Println("2nd item:", ages[1])

	// fmt.Println("2nd item:", ages[2])

	// fmt.Println(ages[len(ages)])
	// here, `len(ages) - 1` equals to 1
	fmt.Println("last item:", ages[len(ages)-1])

	// let's display the indexes and the items
	// of the array
	for i, v := range ages {
		fmt.Printf("index: %d, value: %d\n", i, v)
	}
}
