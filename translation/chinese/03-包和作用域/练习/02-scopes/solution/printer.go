// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

func hello() {
	// only this file can access the imported fmt package
	// when others also do so, they can also access
	//   their own `fmt` "name"

	fmt.Println("hi! this is inanc!")
	bye()
}
