// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"
)

// Since, you didn't learn about the control flow statements yet
// I didn't include an error detection here.
//
// So, if you don't pass a name,
// this program will fail.
// This is intentional.

func main() {
	// assign a new value to the string variable below
	name := os.Args[1]
	fmt.Println("Hello great", name, "!")

	// changes the name, declares the age with 85
	name, age := "gandalf", 85

	fmt.Println("My name is", name)
	fmt.Println("My age is", age)
	fmt.Println("BTW, you shall pass!")
}
