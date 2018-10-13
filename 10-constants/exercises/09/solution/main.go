// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	const (
		Spring = (iota + 1) * 3
		Summer
		Fall
		Winter
	)

	fmt.Println(Winter, Spring, Summer, Fall)
}
