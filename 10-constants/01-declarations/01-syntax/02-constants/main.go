// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// This program uses a named constant
// instead of a magic value

func main() {
	const meters int = 100

	cm := 100
	m := cm / meters // using a named constant
	fmt.Printf("%dcm is %dm\n", cm, m)

	cm = 200
	m = cm / meters // using a named constant
	fmt.Printf("%dcm is %dm\n", cm, m)
}
