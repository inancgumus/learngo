// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	a := [2]string{"jane", "lina"}
	b := [2]string{"jane", "lina"}
	c := [2]string{"mike", "joe"}

	fmt.Println("a==b?", a == b) // equal
	fmt.Println("a==c?", a == c) // not equal

	// cannot compare
	d := [3]string{"jane", "lina"}
	// fmt.Println("c==d?", c == d)

	fmt.Printf("type of c is %T\n", c)
	fmt.Printf("type of d is %T\n", d)
}
