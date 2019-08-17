// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

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
