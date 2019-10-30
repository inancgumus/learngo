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
	// integer types
	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64

	// float types
	var f32 float32
	var f64 float64

	// complex types
	var c64 complex64
	var c128 complex128

	// bool type
	var b bool

	// string types
	var s string
	var r rune  // also a numeric type
	var by byte // also a numeric type

	fmt.Println(
		i, i8, i16, i32, i64,
		f32, f64,
		c64, c128,
		b, r, by,
	)

	// You could do it with Println as well
	fmt.Printf("%q\n", s)
}
