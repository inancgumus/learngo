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
	speed := 100 // int
	force := 2.5 // float64

	// 错误: 无效的操作符
	// ERROR: invalid op
	// speed = speed * force

	// 转换可以是一个破坏性的操作
	// `force` 失去了它的分数部分...

	speed = speed * int(force)

	fmt.Println(speed)
}
