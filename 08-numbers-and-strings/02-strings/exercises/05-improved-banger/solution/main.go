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
	"strings"
	"unicode/utf8"
)

func main() {
	msg := os.Args[1]
	l := utf8.RuneCountInString(msg)

	s := msg + strings.Repeat("!", l)

	fmt.Println(s)
}
