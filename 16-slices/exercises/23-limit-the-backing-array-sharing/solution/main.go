// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"

	"github.com/inancgumus/learngo/16-slices/exercises/23-limit-the-backing-array-sharing/solution/api"
)

func main() {
	temps := api.Read(0, 3)
	temps = append(temps, []int{1, 3}...)

	fmt.Println("API's readings:", api.All())
	fmt.Println("Your readings :", temps)
}
