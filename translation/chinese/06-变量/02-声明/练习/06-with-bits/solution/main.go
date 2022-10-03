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
	// 整数类型
	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64

	// 浮点类型
	var f32 float32
	var f64 float64

	// 复杂类型
	var c64 complex64
	var c128 complex128

	// 布尔型
	var b bool

	// 字符串类型
	var s string
	var r rune  // 同样是数字类型
	var by byte // 同样是数字类型

	fmt.Println(
		i, i8, i16, i32, i64,
		f32, f64,
		c64, c128,
		b, r, by,
	)

	// 你也可以使用 Println 做到
	fmt.Printf("%q\n", s)
}
