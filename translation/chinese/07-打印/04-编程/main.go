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
	var (
		planet   = "venus"
		distance = 261
		orbital  = 224.701
		hasLife  = false
	)

	// 全能打印 %v 动词
	fmt.Printf("Planet: %v\n", planet)
	fmt.Printf("Distance: %v millions kms\n", distance)
	fmt.Printf("Orbital Period: %v days\n", orbital)
	fmt.Printf("Does %v have life? %v\n", planet, hasLife)

	// 参数索引 - 索引从 1 开始
	fmt.Printf(
		"%v is %v away. Think! %[2]v kms! %[1]v OMG.\n",
		planet, distance,
	)

	// 为什么不使用别的打印词?
	// 原因: 有可能因为类型不正确导致报错
	//
	// fmt.Printf("Planet: %d\n", planet)
	// fmt.Printf("Distance: %s millions kms\n", distance)
	// fmt.Printf("Orbital Period: %t days\n", orbital)
	// fmt.Printf("Does %v has life? %f\n", planet, hasLife)

	// 正确的打印动词:
	// fmt.Printf("Planet: %s\n", planet)
	// fmt.Printf("Distance: %d millions kms\n", distance)
	// fmt.Printf("Orbital Period: %f days\n", orbital)
	// fmt.Printf("Does %s has life? %t\n", planet, hasLife)

	// 更加精确的打印动词
	fmt.Printf("Orbital Period: %f days\n", orbital)
	fmt.Printf("Orbital Period: %.0f days\n", orbital)
	fmt.Printf("Orbital Period: %.1f days\n", orbital)
	fmt.Printf("Orbital Period: %.2f days\n", orbital)
}
