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
	fmt.Println(headerOf("gif"))
}

var headers = map[string]string{
	"png": "\x89PNG\r\n\x1a\n",
	"jpg": "\xff\xd8\xff",
}

func headerOf(format string) (header string) {
	defer func() {
		if header == "" {
			panic("unknown format: " + format)
		}
	}()
	return headers[format]
}
