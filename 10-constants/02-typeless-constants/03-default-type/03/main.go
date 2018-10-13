// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	i := 42
	f := 3.14
	b := true
	s := "Hello"
	r := 'A'

	fmt.Printf("i is %T\n", i)
	fmt.Printf("f is %T\n", f)
	fmt.Printf("b is %T\n", b)
	fmt.Printf("s is %T\n", s)

	// int32 = rune (type alias)
	fmt.Printf("r is %T\n", r)
}
