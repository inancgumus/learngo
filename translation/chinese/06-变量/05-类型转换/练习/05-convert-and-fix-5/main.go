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
// 练习: 转换并修复 #5
//
//  修复代码.
//
// 提示
//   int8 的最大值是 127
//   int16 的最大值是 32767
//
// 期望输出
//  1127
// ---------------------------------------------------------

func main() {
	// 不要修改这些变量
	min := int8(127)
	max := int16(1000)

	// 修复代码
	fmt.Println(int8(max) + min)
}
