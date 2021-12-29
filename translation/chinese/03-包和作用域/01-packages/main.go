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

	// 你可以使用其他文件的函数
	// 只要它们在同一个包中

	// 因此, `main()` 可以调用 `bye()` and `hey()`

	// 因为 bye.go, hey.go 和 main.go
	//  都在 main package 中

	bye()
	hey()
}
