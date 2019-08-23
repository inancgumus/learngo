// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

type book struct {
	title string
	price money
}

func (b book) print() {
	fmt.Printf("%-15s: %s\n", b.title, b.price.string())
}
