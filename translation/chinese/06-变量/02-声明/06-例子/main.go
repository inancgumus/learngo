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
	// 变量名是大小写敏感的
	// MyAge, myAge 和 MYAGE 是不同的变量

	// 使用例:
	// 什么时候使用平行声明?
	//
	// 不好:
	// var myAge int
	// var yourAge int
	//
	// 一般:
	// var (
	// 	myAge int
	// 	yourAge int
	// )
	//
	// 推荐:
	var myAge, yourAge int
	fmt.Println(myAge, yourAge)

	var temperature float64
	fmt.Println(temperature)

	var success bool
	fmt.Println(success)

	var language string
	fmt.Println(language)
}
