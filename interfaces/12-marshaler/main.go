// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
)

func main() {
	// removed error handling for brevity (normally you shouldn't do this).
	out, _ := encode()
	fmt.Println(string(out))

	l, _ := decode(out)
	fmt.Print(l)
}
