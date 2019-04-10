// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	words := []string{
		"gopher",
		"programmer",
		"go language",
		"go standard library",
	}

	var bwords [][]byte
	for _, w := range words {
		bw := []byte(w)

		fmt.Println(bw)

		bwords = append(bwords, bw)
	}

	for _, w := range bwords {
		fmt.Println(string(w))
	}
}
