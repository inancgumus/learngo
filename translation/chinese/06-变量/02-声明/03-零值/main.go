// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// 练习: 运行这个, 然后自己查看这些零值 (这里的 零值 指的是变量未初始化时的值, 根据类型的不同会有不同的结果, 而非字面意义上的 "0")
func main() {
	var speed int    // 数字类型
	var heat float64 // 数字类型
	var off bool
	var brand string

	fmt.Println(speed)
	fmt.Println(heat)
	fmt.Println(off)

	// 我用 Printf 来打印空字符串
	// 练习: 用 Println 看看会发生什么
	fmt.Printf("%q\n", brand)
}
