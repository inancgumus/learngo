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
	// 当你在计算中使用一个浮点数和一个整数时
	// 结果总是一个浮点数。

	ratio := 3.0 / 2

	// OR:
	// ratio = 3 / 2.0

	// OR - 如果3在一个int变量里:
	// n := 3
	// ratio = float64(n) / 2

	fmt.Printf("%f", ratio)
}
