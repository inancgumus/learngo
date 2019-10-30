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
	"os"
)

func main() {
	// i := 1
	// fmt.Printf("%q\n", os.Args[1])
	// fmt.Printf("%q\n", os.Args[i])

	// i++
	// fmt.Printf("%q\n", os.Args[2])
	// fmt.Printf("%q\n", os.Args[i])

	// i++
	// fmt.Printf("%q\n", os.Args[3])
	// fmt.Printf("%q\n", os.Args[i])

	// --------------------------------
	// #1st way:
	// --------------------------------

	// for i := 1; i < len(os.Args); i++ {
	// 	fmt.Printf("%q\n", os.Args[i])
	// }

	// --------------------------------
	// #2nd way:
	// --------------------------------

	// for i, v := range os.Args {
	// 	if i == 0 {
	// 		continue
	// 	}
	// 	fmt.Printf("%q\n", v)
	// }

	// --------------------------------
	// #3rd way (best):
	// --------------------------------

	for _, v := range os.Args[1:] {
		fmt.Printf("%q\n", v)
	}
}
