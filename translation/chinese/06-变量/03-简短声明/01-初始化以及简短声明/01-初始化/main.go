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
	// = 是赋值操作符
	// 当在一个变量声明中使用时，它将该变量初始化为给定的值

	// 在这, Go 将 safe 变量初始化为 true

	// 方案 #1（方案 #2 更好）
	// var safe bool = true

	// 方案 #2
	var safe = true

	fmt.Println(safe)
}
