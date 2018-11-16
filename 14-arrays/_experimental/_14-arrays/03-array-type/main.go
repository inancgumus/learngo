// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	ages := [3]int{15, 30, 25}

	fmt.Printf("%T\n", ages)                 // [3]int
	fmt.Printf("%T\n", [...]int{15, 30, 25}) // [3]int

	fmt.Printf("%T\n", [2]int{15, 30}) // [2]int

	fmt.Printf("%T\n", [1]string{"hi"})          // [1]string
	fmt.Printf("%T\n", [...]float64{3.14, 6.28}) // [2]float64
}
