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
	fmt.Println("Hello!")

	// 两个文件属于同一个包
	// 在这里调用 bye.go 的 `bye()`
	bye()
}

// 练习: 删除下面函数的注释
//           分析错误信息

// func bye() {
// 	fmt.Println("Bye!")
// }

// ***** EXPLANATION *****
//
// ERROR: "bye" function "redeclared"
//        in "this block"
//
// "this block" 是指 = "main package"
//
// "redeclared" 是指在同一个作用域中使用了相同的名字
//
// main package's 作用域是:
// main 包中的所有源代码文件
