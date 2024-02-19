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
	n, m := 1, 5

	fmt.Println(2 + 1*m/n)
	fmt.Println(2 + ((1 * m) / n)) // 同上

	// 让我们用括号来改变优先顺序
	fmt.Println(((2 + 1) * m) / n)
}
