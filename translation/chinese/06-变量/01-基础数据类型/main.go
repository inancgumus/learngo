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
	// 整数型
	fmt.Println(
		-200, -100, 0, 50, 100, 100,
	)

	// 浮点型
	fmt.Println(
		-50.5, -20.5, -0., 1., 100.56, // ...
	)

	// 布尔值
	fmt.Println(
		true, false,
	)

	// 字符串 - utf-8
	fmt.Println(
		"", // 空字符串, 只打印一个空格
		"hi",

		// unicode
		"nasılsın?",   // 土耳其语的 "你好吗"
		"hi there 星!", // "你好，star！"
	)
}
