// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	name, last := "carl", "sagan"

	// assignment operation using string concat
	name += " edward"

	// equals to this:
	// name = name + " edward"

	fmt.Println(name + " " + last)
}
