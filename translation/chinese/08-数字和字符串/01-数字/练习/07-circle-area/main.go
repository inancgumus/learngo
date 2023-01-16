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

// ---------------------------------------------------------
// 练习: 圆面积
//
//  根据给定的半径(radius)计算出圆的面积
//
// 圆面积计算公式
//  area = πr²
//  https://en.wikipedia.org/wiki/Area#Circles
//
// 提示
//  对于 PI 你可以使用 `math.Pi`
//
// 预期输出
//  radius: 10 -> area: 314.1592653589793
//
// BONUS 练习!
//  1. 将 area 打印为 314.16
//  2. 要做到这一点，你需要使用正确的Printf动词 :)
//      而不是下面的 `%g` 动词.
//
//    预期输出
//     radius: 10 -> area: 314.16
// ---------------------------------------------------------

func main() {
	var (
		radius = 10.
		area   float64
	)

	// 在这里添加你的代码
	// ...

	// 别碰这个
	fmt.Printf("radius: %g -> area: %g\n", radius, area)
}
