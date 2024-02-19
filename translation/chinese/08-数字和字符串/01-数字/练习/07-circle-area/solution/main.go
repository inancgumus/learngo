// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		radius = 10.
		area   float64
	)

	area = math.Pi * radius * radius

	fmt.Printf("radius: %g -> area: %.2f\n",
		radius, area)

	// 替代方案:
	// math.Pow 计算一个浮点数的幂
	// area = math.Pi * math.Pow(radius, 2)
}
