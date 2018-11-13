// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	msg := "WONDERFUL!"
	bytes := []byte(msg)

	fmt.Println("msg              :", msg)
	fmt.Println("bytes            :", bytes)
	fmt.Println("string(bytes)    :", string(bytes))
	fmt.Println("string(87)       :", string(87))

	fmt.Println()

	for i, v := range msg {
		fmt.Printf(
			"msg[%d]           : %d = %[2]q\n",
			i, v)
	}
}
