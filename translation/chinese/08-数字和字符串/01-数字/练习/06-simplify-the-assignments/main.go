// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// 练习: 简化赋值运算符
//
//  简化代码（重构）
//
// 限制
//  只使用自增自减运算符和赋值运算符
//
// 预期输出
//  3
// ---------------------------------------------------------

func main() {
	width, height := 10, 2

	width = width + 1
	width = width + height
	width = width - 1
	width = width - height
	width = width * 20
	width = width / 25
	width = width % 5

	fmt.Println(width)
}
