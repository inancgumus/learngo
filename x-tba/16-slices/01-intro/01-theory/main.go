// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	// -----------------------------------------------------
	// COMPLEX LITERALS:
	// -----------------------------------------------------
	// array literal : [4]int{1, 2, 3}
	// slice literal : []int{1, 2, 3}
	//
	// You didn't learn about these yet:
	//
	// map literal   : map[int]int{1: 2}
	// struct literal: struct{ x, y int }{ 1, 2 }

	// -----------------------------------------------------
	// SLICE LITERAL:
	// -----------------------------------------------------
	r := []rune{'h', 'e', 'y'} // 3 items
	i := []int{'h', 'e', 'y'}  // 3 items
	b := []byte{'h', 'e', 'y'} // 3 items
	s := []string{"hey"}       // 1 item

	fmt.Printf("[]rune  type: %T\n", r)
	fmt.Printf("[]int   type: %T\n", i)
	fmt.Printf("[]byte  type: %T\n", b)
	fmt.Printf("[]string type: %T\n", s)

	// -----------------------------------------------------
	// multi-dimensional slice
	// -----------------------------------------------------
	multi := [][]rune{
		{1, 2, 3},
		// each item can have a different length
		{4, 5, 6, 7, 8},
	}
	fmt.Printf("multi's type     : %T\n", multi)
	fmt.Printf("multi's length   : %d\n", len(multi))
	fmt.Printf("multi[0]'s length: %d\n", len(multi[0]))
	fmt.Printf("multi[1]'s length: %d\n", len(multi[1]))

	// -----------------------------------------------------
	// getting and setting elements
	// -----------------------------------------------------
	hello := []rune{'h', 'e', 'y'}
	fmt.Println(hello[0], hello[1], hello[2])

	hello[0] = 'y'
	hello[1] = 'o'
	hello[2] = '!'
	fmt.Println(hello[0], hello[1], hello[2])

	// -----------------------------------------------------
	// empty slice
	// -----------------------------------------------------
	e := []rune{}
	fmt.Printf("empty slice's length: %d\n", len(e))

	// -----------------------------------------------------
	// nil slice
	// slice's zero value is nil
	// -----------------------------------------------------
	var n []rune
	fmt.Println("Really nil?", n == nil)

	// you can't get/set elements of a nil slice
	// if you run the code below, it will error at runtime
	//
	// fmt.Println(n[0])
	// fmt.Println(n[1])

	// but you can get its length
	fmt.Printf("nil slice's length: %d\n", len(n))
}
