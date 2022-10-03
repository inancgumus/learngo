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

// NOTE: RUN THIS WITH 3 ARGUMENTS AT LEAST
//       OR, THERE WILL BE AN ERROR

func main() {
	fmt.Printf("%#v\n", os.Args)

	// Gets an item from the os.Args string slice:
	//     os.Args[INDEX]
	// INDEX can be 0 or greater
	fmt.Println("Path:", os.Args[0])
	fmt.Println("1st argument:", os.Args[1])
	fmt.Println("2nd argument:", os.Args[2])
	fmt.Println("3rd argument:", os.Args[3])

	// `len` function can find how many items
	// inside a slice value
	fmt.Println("Items inside os.Args:", len(os.Args))
}
