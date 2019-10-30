// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

type logger struct {
	print   func(...interface{}) (int, error)
	printf  func(fmt string, args ...interface{}) (int, error)
	println func(...interface{}) (int, error)
}

func (l logger) Print(args ...interface{}) {
	l.print(args...)
}

func (l logger) Printf(fmt string, args ...interface{}) {
	l.printf(fmt, args...)
}

func (l logger) Println(args ...interface{}) {
	l.println(args...)
}
