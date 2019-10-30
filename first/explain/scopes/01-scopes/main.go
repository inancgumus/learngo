// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// file scope
import "fmt"

// package scope
const ok = true

// package scope
func main() { // block scope starts

	var hello = "Hello"

	// hello and ok are visible here
	fmt.Println(hello, ok)

} // block scope ends
