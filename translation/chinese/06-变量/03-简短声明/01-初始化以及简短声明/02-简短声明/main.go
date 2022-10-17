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
	// 方案 #1（方案 #2 更好）
	// var safe bool = true

	// 方案 #2 (还行)
	// var safe = true

	// 方案 #3 - 简短声明 (最佳)
	//
	// 你甚至不需要输入 var 关键词
	//
	// 简短声明等同于:
	//   var safe bool = true
	//   var safe = true
	//
	// Go从初始值中获取（推断）类型
	//
	// true 的默认类型是 bool
	// 所以，变量 safe 的类型变成了 bool
	safe := true

	fmt.Println(safe)
}
