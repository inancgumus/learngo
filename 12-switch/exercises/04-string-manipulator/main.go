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
	"strings"
)

// ---------------------------------------------------------
// STORY
//
//	You want to write a program that will manipulate a
//	given string to uppercase, lowercase, and title case.
//
// EXERCISE: String Manipulator
//
//  1. Get the operation as the first argument.
//
//  2. Get the string to be manipulated as the 2nd argument.
//
// HINT
//
//	You can find the manipulation functions in the strings
//	package Go documentation (ToLower, ToUpper, Title).
//
// EXPECTED OUTPUT
//
//	go run main.go
//	  [command] [string]
//
//	  Available commands: lower, upper and title
//
//	go run main.go lower 'OMG!'
//	  omg!
//
//	go run main.go upper 'omg!'
//	  OMG!
//
//	go run main.go title "mr. charles darwin"
//	  Mr. Charles Darwin
//
//	go run main.go genius "mr. charles darwin"
//	  Unknown command: "genius"
//
// ---------------------------------------------------------
const usage = `[command] [string]

Available commands: lower, upper and title
	`

func main() {

	args := os.Args
	if len(args) != 3 {
		fmt.Println(usage)
		return
	}
	command, word := args[1], args[2]
	switch command {
	case "lower":
		word = strings.ToLower(word)
	case "upper":
		word = strings.ToUpper(word)
	case "title":
		word = strings.Title(word)
	default:
		fmt.Printf("Unknown command: %q\n", word)
		return
	}
	fmt.Println(word)
}
