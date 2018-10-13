// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	const min int = 42

	var f float64

	// ERROR: Type Mismatch
	// f = min // NOT OK

	fmt.Println(f)
}
