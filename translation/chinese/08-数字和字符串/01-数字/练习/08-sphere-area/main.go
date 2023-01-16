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
// 练习: 球体面积
//
//  1. 从命令行获取半径
//  2. 将其转换为 float64
//  3. 计算球体的表面积
//
// 球体表面积公式
//  area = 4πr²
//  https://en.wikipedia.org/wiki/Sphere#Surface_area
//
// 限制
//  使用 `math.Pow` 来计算面积
//  阅读文档看看它是如何计算的.
//  https://golang.org/pkg/math/#Pow
//
// 预期输出
//  go run main.go 10
//    1256.64
//
//  go run main.go 54.2
//    36915.47
// ---------------------------------------------------------

func main() {
	var radius, area float64

	// 在这里添加你的代码
	// ...

	// 别碰这个
	fmt.Printf("radius: %g -> area: %.2f\n", radius, area)
}
