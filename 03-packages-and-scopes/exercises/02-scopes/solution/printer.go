// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func hello() {
	// only this file can access the imported fmt package
	// when others also do so, they can also access
	//   their own `fmt` "name"

	fmt.Println("hi! this is inanc!")
	bye()
}
