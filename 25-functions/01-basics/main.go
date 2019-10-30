// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

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
// Because, the compiler needs to see bad.go too
// It can't magically find bad.go — what you give is what you get.
//

package main

// N is a shared counter which is BAD
var N int

func main() {
	showN()
	incrN()
	incrN()
	showN()
}
