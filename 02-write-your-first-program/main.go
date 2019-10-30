// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// package main is a special package
// it allows Go to create an executable file
package main

/*
This is a multi-line comment.

import keyword makes another package available
  for this .go "file".

import "fmt" lets you access fmt package's functionality
  here in this file.
*/
import "fmt"

// "func main" is special.
//
// Go has to know where to start
//
// func main creates a starting point for Go
//
// After compiling the code,
// Go runtime will first run this function
func main() {
	// after: import "fmt"
	// Println function of "fmt" package becomes available

	// Look at what it looks like by typing in the console:
	//   go doc -src fmt Println

	// Println is just an exported function from
	//   "fmt" package

	// Exported = First Letter is uppercase
	fmt.Println("Hello Gopher!")

	// Go cannot call Println function by itself.
	// That's why you need to call it here.
	// It only calls `func main` automatically.

	// -----

	// Go supports Unicode characters in string literals
	// And also in source-code: KÖSTEBEK!
	//
	// Because: Literal ~= Source Code

	// EXERCISE: Remove the comments from below --> //
	// fmt.Println("Merhaba Köstebek!")

	// Unnecessary note:
	// "Merhaba Köstebek" means "Hello Gopher"
	// in Turkish language
}
