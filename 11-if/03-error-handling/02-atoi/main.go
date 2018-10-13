// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Atoi returns an error value
	// So, you should always check it

	n, err := strconv.Atoi(os.Args[1])

	fmt.Println("Converted number    :", n)
	fmt.Println("Returned error value:", err)
}
