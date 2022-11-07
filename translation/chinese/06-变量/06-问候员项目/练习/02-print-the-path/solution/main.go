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

// 步骤:
//
// 输入下面的命令来编译:
//   go build -o myprogram
//
// 然后输入下面的命令来运行:
//   ./myprogram
//
// 如果是 Windows, 输入:
//   myprogram

func main() {
	fmt.Println(os.Args[0])
}
