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
	// 当一个整数和一个浮点数一起使用时
	// 在一个表达式中，其结果总是变成一个浮点数
	fmt.Println(8 * -4.0) // -32.0 not -32

	// 两个整数值的结果是一个整数值
	fmt.Println(-4 / 2)

	// 取余运算符
	// 它只能用于整数
	fmt.Println(5 % 2)
	// fmt.Println(5.0 % 2) // 错误的

	// 加法运算符
	fmt.Println(1 + 2.5)
	fmt.Println(2 - 3)

	// 负数运算符
	fmt.Println(-(-2))
	fmt.Println(- -2) // 这也可以
}
