// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// 正常声明的使用例

// -----------------------------------------------------
// 当你需要一个包作用域变量时
// -----------------------------------------------------

// version := 0 // 不能这样
var version int

func main() {

	// -----------------------------------------------------
	// 如果你不知道初始值时
	// -----------------------------------------------------

	// 别这样做:
	// score := 0

	// 这样做:
	// var score int

	// -----------------------------------------------------
	// 对变量进行分组以提高可读性
	// -----------------------------------------------------

	// var (
	// 	video    string

	// 	duration int
	// 	current  int
	// )
}
