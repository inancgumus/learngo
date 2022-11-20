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
	"os"
)

// 注意: 至少要用3个参数来运行这个程序
//      否则它会报错

func main() {
	fmt.Printf("%#v\n", os.Args)

	// 从 os.Args 字符串切片中获取一个 item
	//     os.Args[INDEX]
	// INDEX 可以是 0 或者更大
	fmt.Println("Path:", os.Args[0])
	fmt.Println("1st argument:", os.Args[1])
	fmt.Println("2nd argument:", os.Args[2])
	fmt.Println("3rd argument:", os.Args[3])

	// `len` 可以知道有多少个 item 在切片值中
	fmt.Println("Items inside os.Args:", len(os.Args))
}
