// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

// REFACTORED VERSION
// It uses well-defined names instead of magic numbers.
// Thanks to keyed elements and constants.

func main() {
	const (
		ETH = iota
		WAN
	)

	rates := [...]float64{
		ETH: 25.5,
		WAN: 120.5,
	}

	// uses well-defined names (ETH, WAN, ...) - good
	fmt.Printf("1 BTC is %g ETH\n", rates[ETH])
	fmt.Printf("1 BTC is %g WAN\n", rates[WAN])
}
