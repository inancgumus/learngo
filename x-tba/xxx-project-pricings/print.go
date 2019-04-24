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

// show shows the given properties
func show(props []property) {
	showHeader()

	for _, p := range props {
		fmt.Printf("%-15s", p.location)
		fmt.Printf("%-15d", p.size)
		fmt.Printf("%-15d", p.beds)
		fmt.Printf("%-15d", p.baths)
		fmt.Printf("%-15d", p.price)
		fmt.Println()
	}
}

// showHeader prints the header
func showHeader() {
	const header = "Location,Size,Beds,Baths,Price"

	for _, h := range strings.Split(header, separator) {
		fmt.Printf("%-15s", h)
	}
	fmt.Printf("\n%s\n", strings.Repeat("=", 75))
}
