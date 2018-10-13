// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	// this one uses a raw string literal

	// can you see how readable it is?
	// compared to the previous one?

	path := `c:\program files\duper super\fun.txt
c:\program files\really\funny.png`

	fmt.Println(path)
}
