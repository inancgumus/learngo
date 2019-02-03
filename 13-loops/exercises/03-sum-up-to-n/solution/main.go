// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("gimme two numbers")
		return
	}

	min, err1 := strconv.Atoi(os.Args[1])
	max, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil || err2 != nil || min >= max {
		fmt.Println("wrong numbers")
		return
	}

	var sum int
	for i := min; i <= max; i++ {
		sum += i

		fmt.Print(i)
		if i < max {
			fmt.Print(" + ")
		}
	}
	fmt.Printf(" = %d\n", sum)
}
