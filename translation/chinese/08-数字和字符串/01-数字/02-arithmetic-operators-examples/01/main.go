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
	fmt.Println("sum :", 3+2)   // sum - int
	fmt.Println("sum :", 2+3.5) // sum - float64
	fmt.Println("dif :", 3-1)   // difference - int
	fmt.Println("dif :", 3-0.5) // difference - float64
	fmt.Println("prod:", 4*5)   // product - int
	fmt.Println("prod:", 5*2.5) // product - float64
	fmt.Println("quot:", 8/2)   // quotient - int
	fmt.Println("quot:", 8/1.5) // quotient - float64

	// 余数仅适用于整数
	fmt.Println("rem :", 8%3)
	// fmt.Println("rem:", 8.0%3) // error

	// 你可以这样做
	// 因为浮点数的小数部分为零
	f := 8.0
	fmt.Println("rem :", int(f)%3)
}
