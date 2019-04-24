// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

//
// You need run this program like so:
//   go run .
//
// This will magically pass all the go files in the current directory to the
// Go compiler.
//
//
// BUT NOT like so:
//   go run main.go
//
// Because, the compiler needs to see global.go too
// It can't magically find global.go — what you give is what you get.
//

package main

import (
	"fmt"
	"strconv"
)

func main() {
	local := 10
	show(local)

	incrWrong(local)
	fmt.Printf("local → %d\n", local)

	// incr(local)
	local = incr(local)
	show(local)

	local = incrBy(local, 5)
	show(local)

	_, err := incrByStr(local, "TWO")
	if err != nil {
		fmt.Printf("err   → %s\n", err)
	}

	local, _ = incrByStr(local, "2")
	show(local)

	// CHAINING

	// can't save the number back into main's local
	show(incrBy(local, 2))
	show(local)

	// can't pass the results of the multiple-value returning func
	// show(incrByStr(local, "3"))

	// can call showErr directly because it accepts the same types
	// of parameters as with incrByStr's result values.
	// local = sanitize(incrByStr(local, "NOPE"))
	// show(local)

	local = sanitize(incrByStr(local, "2"))
	show(local)

	local = limit(incrBy(local, 5), 2000)
	show(local)
}

func show(n int) {
	// can't access main's local
	// fmt.Printf("show : n = %d\n", local)
	fmt.Printf("show  → n = %d\n", n)
}

// wrong: incr can't access to main's `local`
func incrWrong(n int) {
	// n := local
	n++
	// can't return (there are no result values)
	// return n
}

func incr(n int) int {
	n++
	return n
}

func incrBy(n, factor int) int {
	return n * factor
}

func incrByStr(n int, factor string) (int, error) {
	m, err := strconv.Atoi(factor)
	n = incrBy(n, m)
	return n, err
}

func sanitize(n int, err error) int {
	if err != nil {
		return 0
	}
	return n
}

func limit(n, lim int) (m int) {
	// var m int
	m = n
	if m >= lim {
		m = lim
	}
	// return m
	return
}
