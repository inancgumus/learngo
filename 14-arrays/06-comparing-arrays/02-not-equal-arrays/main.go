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
	b := [3]int{3, 9, 6}

	if a == b {
		fmt.Println("they're equal!")
	} else {
		fmt.Println("they're not equal! a==b?", a == b)

		// but, they're comparable
		// since their types are identical
		fmt.Printf("1st array's type is %T\n", a)
		fmt.Printf("2nd array's type is %T\n", b)

		// go compares arrays like this
		// it compares the elements one by one
		fmt.Println(a[0] == b[0]) // false
		fmt.Println(a[1] == b[1]) // true
		fmt.Println(a[2] == b[2]) // false

		// actually Go will quit comparing these arrays
		//   once it sees unequal items
		//
		// so, it will stop the comparison (short-circuit)
		//   when it sees the first item is false
		//   and it will return false
	}
}
