// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	name, lastname := "Nikola", "Tesla"
	fmt.Println(name, lastname)

	birth, death := 1856, 1943
	fmt.Println(birth, death)

	on, off := true, false
	fmt.Println(on, off)

	// there's no limit
	// however, more declarations that you declare
	// more unreadable it becomes...
	degree, ratio, heat := 10.55, 30.5, 20.
	fmt.Println(degree, ratio, heat)

	// you can short declare variables with different types
	nFiles, valid, msg := 10, true, "hello"

	fmt.Println(nFiles, valid, msg)
}
