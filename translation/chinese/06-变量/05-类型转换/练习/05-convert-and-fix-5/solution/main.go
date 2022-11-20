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
	min := int8(127)
	max := int16(1000)

	fmt.Println(max + int16(min))

	// 解释
	//
	// `int8(max)` 损失了 max 的信息
	// 导致它减小到了 127
	// 也就是 int8 的最大值
	//
	// 正确的转换时 int16(min)
	// 因为 int16 > int8
	// 当你这样做时, min 不会损失信息
	//
	// 你会在 "Go 类型系统" 这节学到更多

}
