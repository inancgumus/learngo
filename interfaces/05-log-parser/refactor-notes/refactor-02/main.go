// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
)

func main() {
	res, err := parse()
	if err != nil {
		fmt.Println("> Err:", err)
		return
	}

	summarize(res)
}
