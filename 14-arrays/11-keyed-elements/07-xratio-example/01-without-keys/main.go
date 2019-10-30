// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

func main() {
	rates := [...]float64{
		25.5,  // ethereum
		120.5, // wanchain
	}

	// uses magic values - not good
	fmt.Printf("1 BTC is %g ETH\n", rates[0])
	fmt.Printf("1 BTC is %g WAN\n", rates[1])
}
