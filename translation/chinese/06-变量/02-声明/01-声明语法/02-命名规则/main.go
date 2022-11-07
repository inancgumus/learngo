// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// 变量声明规则

func main() {
	// 正确的声明
	var speed int
	var SpeeD int

	// 下划线开头是允许但不推荐的
	var _speed int

	// Unicode 字母是允许的
	var 速度 int

	// 让编译器开心
	_, _, _, _ = speed, SpeeD, _speed, 速度
}
