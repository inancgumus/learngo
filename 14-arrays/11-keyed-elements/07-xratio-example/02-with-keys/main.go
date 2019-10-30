// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// REFACTORED VERSION
// It uses well-defined names instead of magic numbers.
// Thanks to the keyed elements and constants.

func main() {
	const (
		ETH = 9 - iota
		WAN
		ICX
		// you can add more cryptocurrencies here
		// watch out the -1 index though!
	)

	rates := [...]float64{
		ETH: 25.5,
		WAN: 120.5,
		ICX: 20,
		// you can add more cryptocurrencies here
	}

	// uses well-defined names (ETH, WAN, ...) - good
	fmt.Printf("1 BTC is %g ETH\n", rates[ETH])
	fmt.Printf("1 BTC is %g WAN\n", rates[WAN])
	fmt.Printf("1 BTC is %g ICX\n", rates[ICX])

	fmt.Printf("%#v\n", rates)
}
