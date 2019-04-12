// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package pricing

import (
	"fmt"
	"strings"
)

// Print prints the given properties
func Print(props Properties) {
	printHeader()

	for _, p := range props.list {
		fmt.Printf("%-15s", p.Location)
		fmt.Printf("%-15d", p.Size)
		fmt.Printf("%-15d", p.Beds)
		fmt.Printf("%-15d", p.Baths)
		fmt.Printf("%-15d", p.Price)
		fmt.Println()
	}
}

// printHeader prints the header
func printHeader() {
	const header = "Location,Size,Beds,Baths,Price"

	for _, h := range strings.Split(header, separator) {
		fmt.Printf("%-15s", h)
	}
	fmt.Printf("\n%s\n", strings.Repeat("=", 75))
}
