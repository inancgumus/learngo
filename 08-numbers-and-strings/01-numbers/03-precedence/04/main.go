// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	celsius := 35.

	// Wrong formula  :  9*celsius + 160  / 5
	// Correct formula: (9*celsius + 160) / 5
	fahrenheit := (9*celsius + 160) / 5

	fmt.Printf("%g ºC is %g ºF\n", celsius, fahrenheit)
}
