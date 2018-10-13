// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	fmt.Println("Hello!")

	// two files belong to the same package
	// calling `bye()` of bye.go here
	bye()
}

// EXERCISE: Remove the comments from the below function
//           And analyze the error message

// func bye() {
// 	fmt.Println("Bye!")
// }

// ***** EXPLANATION *****
//
// ERROR: "bye" function "redeclared"
//        in "this block"
//
// "this block" means = "main package"
//
// "redeclared" means = you're using the same name
//   in the same scope again
//
// main package's scope is:
// all source-code files which are in the same main package
