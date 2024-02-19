// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// 自增自减运算符是一个语句

func main() {
	var counter int

	// 以下的 "语句" 是正确的:

	counter++ // 1
	counter++ // 2
	counter++ // 3
	counter-- // 2
	fmt.Printf("There are %d line(s) in the file\n",
		counter)

	// 以下的 "表达式" 是不正确的:

	// counter = 5+counter--
	// counter = ++counter + counter--
}
