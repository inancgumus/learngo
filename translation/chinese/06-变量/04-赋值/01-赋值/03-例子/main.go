// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

func main() {
	var (
		name   string
		age    int
		famous bool
	)

	// Example #1
	name = "Newton"
	age = 84
	famous = true

	fmt.Println(name, age, famous)

	// Example #2
	name = "Somebody"
	age = 20
	famous = false

	fmt.Println(name, age, famous)

	// Example #3
	// EXERCISE: Why this doesn't work? Think about it.

	// name = 20
	// name = age

	// save the previous value of the variable
	// to a new variable
	var prevName string
	prevName = name

	// overwrite the value of the original variable
	// by assigning to it
	name = "Einstein"

	fmt.Println("previous name:", prevName)
	fmt.Println("current name :", name)
}
