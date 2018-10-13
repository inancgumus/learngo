// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	a := [3]int{6, 9, 3}
	b := [3]int{6, 9, 3}

	if a == b {
		fmt.Println("they're equal!")

		// they're comparable
		// since their types are identical
		fmt.Printf("1st array's type is %T\n", a)
		fmt.Printf("2nd array's type is %T\n", b)

		// go compares arrays like this
		// it compares the elements one by one
		fmt.Println(a[0] == b[0])
		fmt.Println(a[1] == b[1])
		fmt.Println(a[2] == b[2])
	}
}
