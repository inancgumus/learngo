// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	var brand string

	// prints the string in quoted-form like this ""
	fmt.Printf("%q\n", brand)

	brand = "Google"
	fmt.Printf("%q\n", brand)
}
