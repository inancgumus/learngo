// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("HasPrefix: %T\n", strings.HasPrefix)
	fmt.Printf("HasSuffix: %T\n", strings.HasSuffix)
	fmt.Println()

	var fn func(string, string) bool

	fn = strings.HasPrefix
	fn = strings.HasSuffix

	ok := fn("gopher", "go")

	fmt.Printf("ok       : %t\n", ok)
}
