// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}

	// Store a slice in `any`
	var any interface{}
	any = nums

	// `any` is an empty interface value, it's not a slice.
	// You cannot get its length. It has no length.
	// _ = len(any)

	// When you extract the slice from it, you can get the slice's length.
	_ = len(any.([]int))

	// You cannot assign an `[]T` slice to `[]interface{}` slice.
	// Where `T` can be of any type.
	// Reason: `[]interface{}` is a slice, it's not an `empty interface` value.
	var many []interface{}
	// many = nums

	// The same reason that you can't assign an `[]int` to a `[]string`.
	// var words []string = nums

	// You need to copy the elements from `nums` to `many` manually.
	for _, n := range nums {
		many = append(many, n)
	}
	fmt.Println(many)

	// Each element of the `many` slice is an `empty interface` value.
	// el := many[0]
}
