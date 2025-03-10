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
	// Go编译器将这些数字视为整数,
	// 因为整数值中没有小数部分
	// 所以, 结果是1, 而不是1.5

	// 所以, 这里的ratio变量是一个int变量
	// 这是因为, 3除以2的结果是一个整数

	ratio := 3 / 2

	fmt.Printf("%d", ratio)
}
