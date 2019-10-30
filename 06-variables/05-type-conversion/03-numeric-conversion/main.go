// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

func main() {
	var apple int
	var orange int32

	// ERROR:
	// cannot assign orange to apple (different types)
	// apple = orange

	// you need to convert orange to apple

	// orange is convertable to int because,
	//   int and int32 are both numeric types

	apple = int(orange)

	// you can't convert a numeric type to a bool:
	// isDelicious := bool(orange)

	// but you can convert an int to a string
	// this only works with int types
	orange = 65 // 65 is A
	color := string(orange)
	fmt.Println(color)

	// this doesn't work. 65.0 is a float.
	// fmt.Println(string(65.0))

	// this works: there are two byte values
	// byte is also an int
	fmt.Println(string([]byte{104, 105}))

	_ = apple
}
