// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// package main 是一个特殊的包
// 它允许 Go 创建一个可执行文件
package main

/*
这是一个多行注释

import 关键字引入另一个包
  对于这个 .go ”文件“

import "fmt" 允许你使用 fmt 包的各种函数
  在这个文件
*/
import "fmt"

// "func main" 是特殊的
//
// Go 必须知道从哪里开始执行
//
// func main 对于 Go 而言就是开始执行的位置
//
// 代码编译后，
// Go runtime 第一个执行该函数
func main() {
	// import "fmt" 之后
	// "fmt" 包中的 Println 函数就可以使用

	// 在控制台输入下面的命令可以查看它的源代码:
	//   go doc -src fmt Println

	// Println 是一个 exported 函数来自
	//   "fmt" 包

	// Exported = 首字母大写
	fmt.Println("Hello Gopher!")

	// Go 本身不能调用 Println 函数
	// 这就是为什么你需要在这里调用它
	// Go 只能自动地调用 `func main`

	// -----

	// Go 字符串支持 Unicode 字符
	// 源代码同时也支持: KÖSTEBEK!
	//
	// 因为: 字面值 ~= 源代码

	// 练习: 删除下行注释的 --> //
	// fmt.Println("Merhaba Köstebek!")

	// 不重要的注释:
	// "Merhaba Köstebek" 的意思是 "你好 Gopher"
	// 在土耳其语
}
