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
	"time"
)

func main() {
	// #1
	t := time.Second * 10

	fmt.Println(t)

	// #2
	i := 10

	// NOT OK
	// time.Duration and int are different types
	// t = time.Second * i

	// OK: i is int, Duration is int64
	//     So, i is convertable to int64
	t = time.Second * time.Duration(i)

	fmt.Println(t)

	// #3
	const c = 10

	// OK: Because, c is typeless
	t = time.Second * c

	fmt.Println(t)
}
