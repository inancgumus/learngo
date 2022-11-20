// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"
)

// 因为, 还没有学到关于控制流语句的知识
// 在这里没有包含错误检测
//
// 所以, 如果你没有传递 name,
// 程序会出错,
// 这是有意为之的

func main() {
	var name string

	// 给下面的字符串变量赋一个新值
	name = os.Args[1]
	fmt.Println("Hello great", name, "!")

	// 修改 name, 声明 age 并赋值 85
	name, age := "gandalf", 85

	fmt.Println("My name is", name)
	fmt.Println("My age is", age)
	fmt.Println("BTW, you shall pass!")
}
