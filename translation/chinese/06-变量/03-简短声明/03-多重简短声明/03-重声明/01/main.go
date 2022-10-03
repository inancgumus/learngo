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
	// `safe` 的值是 `false`
	var safe bool

	// `safe` 的值现在是 `true`

	// `speed` 被声明并初始化成了 `50`

	// 只在以下情况下能进行重声明
	//
	// 至少一个重声明的变量是新变量

	safe, speed := true, 50

	fmt.Println(safe, speed)

	// 练习
	//
	// 在"再次"简短声明前声明变量 speed ( 即在重声明语句前声明 speed 变量 )
	//
	// Observe what happens
}
