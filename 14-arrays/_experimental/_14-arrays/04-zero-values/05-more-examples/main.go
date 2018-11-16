// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	names := [5]string{
		"lina", "bob", "jane",
	}

	fmt.Printf("%#v\n", names)

	temperatures := [10]float64{1.5, 2.5}
	fmt.Printf("%#v\n", temperatures)
}
