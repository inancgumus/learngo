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

// ---------------------------------------------------------
// 这是我在讲座中提到的练习的答案
// ---------------------------------------------------------

// 注意: 你至少要用 3 个参数来运行这个程序
//       否则它会报错

func main() {
	// 给下面的字符串变量赋一个新值
	var (
		name  = os.Args[1]
		name2 = os.Args[2]
		name3 = os.Args[3]
	)

	fmt.Println("Hello great", name, "!")
	fmt.Println("And hellooo to you magnificient", name2, "!")
	fmt.Println("Welcome", name3, "!")

	// 修改 name, 声明 age 变量并赋值 131
	name, age := "bilbo baggins", 131 // 未知年龄!

	fmt.Println("My name is", name)
	fmt.Println("My age is", age)
	fmt.Println("And, I love adventures!")

}
