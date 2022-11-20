// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// 转化的顺序很重要...

func main() {
	speed := 100
	force := 2.5

	fmt.Printf("speed     : %T\n", speed)
	fmt.Printf("conversion: %T\n", float64(speed))
	fmt.Printf("expression: %T\n", float64(speed)*force)

	// 类型不匹配:
	//   speed 是 int
	//   expression 是 float64
	// speed = float64(speed) * force

	// 正确: int * int
	speed = int(float64(speed) * force)

	fmt.Println(speed)
}
