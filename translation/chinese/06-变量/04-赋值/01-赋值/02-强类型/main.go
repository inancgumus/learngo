// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// Go 是一个强类型的编程语言

// 即使是 浮点 (float) 和 整数 (int) 也是不同的类型
// 甚至: int32 和 int 也是不同的类型

// 练习: 尝试去掉注释并观察错误

func main() {
	var speed int
	// speed = "100"

	var running bool
	// running = 1

	var force float64
	// speed = force

	// Go 自动将无标注的 int 转换为 float64
	force = 1

	// 让编译器开心
	_, _, _ = speed, running, force
}
