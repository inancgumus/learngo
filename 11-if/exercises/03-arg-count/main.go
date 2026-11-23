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

// ---------------------------------------------------------
// EXERCISE: Arg Count
//
//  1. Get arguments from command-line.
//  2. Print the expected outputs below depending on the number
//     of arguments.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me args
//
//  go run main.go hello
//    There is one: "hello"
//
//  go run main.go hi there
//    There are two: "hi there"
//
//  go run main.go I wanna be a gopher
//    There are 5 arguments
// ---------------------------------------------------------

func main() {
	user_args_count := len(os.Args) - 1

	if user_args_count == 0 {
		fmt.Println("Give me args")
	} else if user_args_count == 1 {
		fmt.Printf("There is one: %q\n", os.Args[1])
	} else if user_args_count == 2 {
		fmt.Printf("There are two: %q\n", os.Args[1]+" "+os.Args[2])
	} else {
		fmt.Printf("There are %d arguments\n", user_args_count)
	}
}
