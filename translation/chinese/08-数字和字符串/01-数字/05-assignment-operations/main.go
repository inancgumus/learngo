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
)

func main() {
	width, height := 5., 12.

	// 计算长方形的面积
	area := width * height
	fmt.Printf("%gx%g=%g\n", width, height, area)

	area = area - 10 // area减10
	area = area + 10 // area加10
	area = area * 2  // area乘以2
	area = area / 2  // area除以2
	fmt.Printf("area=%g\n", area)

	// // 赋值运算符
	area -= 10 // area减10
	area += 10 // area加10
	area *= 2  // area乘以2
	area /= 2  // area除以2
	fmt.Printf("area=%g\n", area)

	// 找出 area 变量的余数
	// 因为: area 是浮点数, 所以这不起作用:
	// area %= 7

	// 这种方法可以运作
	area = float64(int(area) % 7)
	fmt.Printf("area=%g\n", area)
}
