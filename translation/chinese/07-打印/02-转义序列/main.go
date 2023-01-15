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
	// 没有换行符
	fmt.Println("hihi")

	// 使用换行符:
	//   \n = 转义序列
	//   \  = 转义字符
	fmt.Println("hi\nhi")

	// 转义字符:
	//   \\ = \
	//   \" = "
	fmt.Println("hi\\n\"hi\"")
}
