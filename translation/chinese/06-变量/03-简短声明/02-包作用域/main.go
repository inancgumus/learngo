// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// 你不能使用没有关键字的声明语句
// 简短声明没有关键字 (`var`)
// 所以, 它不能在包作用域上使用
//
// 语法错误:
// "函数体外的非声明语句"
// 原文:
// SYNTAX ERROR:
// "non-declaration statement outside function body"

// safe := true

// 然而, 你可以在包作用域上正常声明一个变量, 因为它有关键词: `var`
var safe = true

func main() {
	fmt.Println(safe)
}
