// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// 文件作用域
import "fmt"

// 包作用域
const ok = true

// 包作用域
func main() { // 块作用域开始

	var hello = "Hello"

	// hello 和 ok 都可见
	fmt.Println(hello, ok)

} // 块作用域结束
