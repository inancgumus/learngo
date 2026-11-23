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
// 练习: 优先级
//
//  改变表达式以产生预期输出
//
// 限制
//  使用括号来改变优先顺序
// ---------------------------------------------------------

func main() {
	// 这个表达式应该打印 20
	fmt.Println(10 + 5 - 5 - 10)

	// 这个表达式应该打印 -16
	fmt.Println(-10 + 0.5 - 1 + 5.5)

	// 这个表达式应该打印 -25
	fmt.Println(5 + 10*2 - 5)

	// 这个表达式应该打印 0.5
	fmt.Println(0.5*2 - 1)

	// 这个表达式应该打印 24
	fmt.Println(3 + 1/2*10 + 4)

	// 这个表达式应该打印 15
	fmt.Println(10 / 2 * 10 % 7)

	// 这个表达式应该打印 40
	// 请注意，你可能需要使用浮点数来解决这个问题
	fmt.Println(100 / 5 / 2)
}
