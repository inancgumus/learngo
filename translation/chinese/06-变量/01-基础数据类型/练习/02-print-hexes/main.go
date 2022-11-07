// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// 这个练习是可选的

// ---------------------------------------------------------
// 练习: 打印十六进制
//
//  1. 打印十六进制的 0 到 9
//
//  2. 打印十六进制的 10 到 15
//
//  3. 打印十六进制的 17
//
//  4. 打印十六进制的 25
//
//  5. P打印十六进制的 50
//
//  6. 打印十六进制的 100
//
// 期望输出
//  0 1 2 3 4 5 6 7 8 9
//  10 11 12 13 14 15
//  17
//  25
//  50
//  100
//
// NOTE
//  https://stackoverflow.com/questions/910309/how-to-turn-hexadecimal-into-decimal-using-brain
//
// https://simple.wikipedia.org/wiki/Hexadecimal_numeral_system
//
// ---------------------------------------------------------

func main() {
	// 例子

	// 我要打印十六进制的 10
	fmt.Println(0xa)

	// 我要打印十六进制的 16
	// 0x10
	//   ^^-----  1 * 0 = 0
	//   |
	//   +------ 16 * 1 = 16
	//                  = 16
	fmt.Println(0x10)

	// 我要打印十六进制的 150
	// 0x96
	//   ^^-----  1 * 6 = 6
	//   |
	//   +------ 16 * 9 = 144
	//                  = 150
	fmt.Println(0x96)

	// 把上面的所有代码注释掉, 然后
	// 在下面写下你的代码
}
