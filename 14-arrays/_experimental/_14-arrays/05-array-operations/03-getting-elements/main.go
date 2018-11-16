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

	ages[0] = 6
	ages[1] *= 3

	fmt.Println(ages)

	// increase the last element by 1
	ages[1]++
	ages[len(ages)-1]++

	fmt.Println(ages)

	// v is a copy
	for _, v := range ages {
		// it's like:
		// v = ages[0]
		// v++

		// and:
		// v = ages[1]
		// v++

		v++
	}

	fmt.Println(ages)

	// you don't need to use blank-identifier for the value
	// for i, _ := range ages {

	for i := range ages {
		ages[i]++
	}

	fmt.Println(ages)
}
