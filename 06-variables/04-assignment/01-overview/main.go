// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	var counter int

	fmt.Println("counter's name : counter")
	fmt.Println("counter's value:", counter)
	fmt.Printf("counter's type : %T\n", counter)

	counter = 10 // OK
	// counter = "ten" // NOT OK

	fmt.Println("counter's value:", counter)

	counter++
	fmt.Println("counter's value:", counter)
}
