// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

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
