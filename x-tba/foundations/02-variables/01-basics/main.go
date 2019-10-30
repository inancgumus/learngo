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

func main() {
	// var (
	// 	name string
	// 	name2 string
	// 	name3 string
	// )

	// var name, name2, name3 string
	// name = os.Args[1]
	// name2 = os.Args[2]
	// name3 = os.Args[3]

	// name := os.Args[1]
	// name2 := os.Args[2]
	// name3 := os.Args[3]

	name, name2, name3 := os.Args[1], os.Args[2], os.Args[3]

	fmt.Println("Hello great", name, "!")
	fmt.Println("And hellooo to you magnificient", name2, "!")
	fmt.Println("Welcome", name3, "!")

	// changes the name, declares the age with 131
	name, age := "bilbo baggins", 131 // unknown age!

	fmt.Println("My name is", name)
	fmt.Println("My age is", age)
	fmt.Println("And, I love adventures!")
}
