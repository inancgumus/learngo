// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"time"
)

func main() {
	// hours's type is time.Duration
	// but later's type was int
	// now, `later` is typeless
	//
	// time.Duration's underlying type is int64
	// and, `later` is numeric typeless value
	// so, they can be multiplied
	const later = 10

	hours, _ := time.ParseDuration("1h")

	fmt.Printf("%s later...\n", hours*later)
}
