// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// 我不想在这节课中谈论这个话题
// 作为旁注，我想把它放在这里
// 记得复习

var declareMeAgain = 10

func nested() { // 块作用域开始

	// 定义相同的变量
	// 它们可以同时存在
	// 这个只属于块作用域
	// 包作用域的变量仍然可以使用
	var declareMeAgain = 5
	fmt.Println("inside nested:", declareMeAgain)

} // 块作用域结束

func main() { // 块作用域开始

	fmt.Println("inside main:", declareMeAgain)

	nested()

	// 包作用域的 declareMeAgain 不受 nested 函数的影响
	fmt.Println("inside main:", declareMeAgain)

} // 块作用域结束
