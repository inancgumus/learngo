// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	// instead of storing the names in variables,
	// you can use an array to store all of the names
	//   in a single variable

	// names := [3]string{"jake", "joe", "lee"}
	names := [4]string{"jake", "joe", "lee", "lina"}

	// doing so allows you to loop over the names
	// and print each one of them
	//
	// you can't do this without arrays
	for _, name := range names {
		fmt.Println(name)
	}
}
