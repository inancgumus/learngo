// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	counter, factor := 45, 0.5

	counter++
	counter++
	counter++
	counter++
	counter++

	factor--
	factor--

	fmt.Println(float64(counter) * factor)
}
