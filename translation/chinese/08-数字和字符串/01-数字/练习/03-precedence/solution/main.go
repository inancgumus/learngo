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
	// 10 + 5 - 5 - 10
	fmt.Println(10 + 5 - (5 - 10))

	// -10 + 0.5 - 1 + 5.5
	fmt.Println(-10 + 0.5 - (1 + 5.5))

	// 5 + 10*2 - 5
	fmt.Println(5 + 10*(2-5))

	// 0.5*2 - 1
	fmt.Println(0.5 * (2 - 1))

	// 3 + 1/2*10 + 4
	fmt.Println((3+1)/2*10 + 4)

	// 10 / 2 * 10 % 7
	fmt.Println(10 / 2 * (10 % 7))

	// 100 / 5 / 2
	// 5  / 2 = 2
	//  5 和 2 是整数, 所以, 刪除小数部分
	// 5. / 2 = 2.5
	//  因为 5 是一个浮点数，所以结果变成了一个浮点数
	fmt.Println(100 / (5. / 2))
}
