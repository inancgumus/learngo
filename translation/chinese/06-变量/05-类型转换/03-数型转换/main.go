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
	var apple int
	var orange int32

	// 错误:
	// 不能将 orange 赋值到 apple 上 (不同的类型)
	// ERROR:
	// cannot assign orange to apple (different types)
	// apple = orange

	// 你需要将 orange 转换为 apple

	// orange 能够被转换为 int ,
	//   因为 int 和 int32 都是数字类型

	apple = int(orange)

	// 你不能将一个数字类型转换为布尔类型
	// isDelicious := bool(orange)

	// 但是你可以将 int 转换为 string
	// 这只对 int 类型生效
	orange = 65 // 65 is A
	color := string(orange)
	fmt.Println(color)

	// 这不能运行。 65.0 是 float.
	// fmt.Println(string(65.0))

	// 这能运行: 这有两个 byte 值
	// byte 也是 int
	fmt.Println(string([]byte{104, 105}))

	_ = apple

}
