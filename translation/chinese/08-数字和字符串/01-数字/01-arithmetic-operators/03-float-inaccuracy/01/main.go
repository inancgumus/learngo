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
	ratio := 1.0 / 10.0

	// 经过10次操作后
	// 不准确的地方很明显
	//
	// 不过,现在不要介意这个循环的语法
	// 我将在事后解释
	for range [...]int{10: 0} {
		ratio += 1.0 / 10.0
	}

	fmt.Printf("%.60f", ratio)
}
