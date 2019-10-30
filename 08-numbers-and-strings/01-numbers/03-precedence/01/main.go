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
	fmt.Println(
		2+2*4/2,
		2+((2*4)/2), // same as above
	)

	fmt.Println(
		1+4-2,
		(1+4)-2, // same as above
	)

	fmt.Println(
		(2+2)*4/2,
		(2+2)*(4/2), // same as above
	)
}
