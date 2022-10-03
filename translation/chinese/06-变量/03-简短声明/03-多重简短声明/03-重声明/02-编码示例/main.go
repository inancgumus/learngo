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
	// 例 #1

	name := "Nikola"
	fmt.Println(name)

	// name 已经在代码块里存在了
	// name := "Marie"

	// 给 name 赋了一个新值
	// 并且声明一个新变量 age 并赋值 66
	name, age := "Marie", 66
	fmt.Println(name, age)

	// 例 #2

	// name = "Albert"
	// birth := 1879

	// 下面的重声明等同于上面的语句.
	//
	// `name` 是一个已经存在的变量
	//   Go 给变量 name 赋值了 "Albert"
	//
	// `birth` 是一个新变量
	//   Go 声明并给它赋值了 `1879`
	name, birth := "Albert", 1879

	fmt.Println(name, birth)
}
