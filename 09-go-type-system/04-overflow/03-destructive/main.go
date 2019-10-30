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
	// uint16 max value is 65535
	big := uint16(65535)

	// uint8 destroys its value
	// to its own max value which is 255
	//
	// 65535 - 255 is lost.
	small := uint8(big)

	// fmt.Printf("%b %d\n", big, big)
	// fmt.Printf("%b %[1]d\n", big)

	fmt.Printf("%016b %[1]d\n", big)
	fmt.Printf("%016b %[1]d\n", small)
}
