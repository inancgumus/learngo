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
)

func main() {
	fmt.Println("Hello!")

	// 语句改变执行顺序
	// 特别是流程控制语句，如 `if`
	if 5 > 1 {
		fmt.Println("bigger")
	}
}
