// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// 简短声明的使用例

func main() {

	// -----------------------------------------------------
	// 如果你知道初始值
	// -----------------------------------------------------

	// 别这样做:
	// var width, height = 100, 50

	// 这样做 (简洁):
	// width, height := 100, 50

	// -----------------------------------------------------
	// 重声明
	// -----------------------------------------------------

	// 别这样做:
	// width = 50
	// color := red

	// 这样做 (简洁):
	// width, color := 50, "red"
}
