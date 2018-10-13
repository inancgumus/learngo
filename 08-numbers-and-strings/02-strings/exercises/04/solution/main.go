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
	"unicode/utf8"
)

func main() {
	length := utf8.RuneCountInString(os.Args[1])
	fmt.Println(length)
}
