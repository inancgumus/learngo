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
)

// STEPS:
//
// Compile it by typing:
//   go build -o myprogram
//
// Then run it by typing:
//   ./myprogram
//
// If you're on Windows, then type:
//   myprogram

func main() {
	fmt.Println(os.Args[0])
}
