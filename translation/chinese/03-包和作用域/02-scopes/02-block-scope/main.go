// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

func nope() { //块作用域开始

	// hello 和 ok 只定义在当前块作用域
	const ok = true
	var hello = "Hello"

	_ = hello
} // 块作用域结束

func main() { // 块作用域开始

	// hello 和 ok 不可见

	// ERROR:
	// fmt.Println(hello, ok)

} // 块作用域结束
