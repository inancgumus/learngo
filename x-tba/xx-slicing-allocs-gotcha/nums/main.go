// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	const path = "nums.txt"

	_ = os.Remove(path)

	f, _ := os.OpenFile(path,
		os.O_CREATE|os.O_WRONLY,
		0644)
	defer f.Close()

	const size = 1000000
	w := bufio.NewWriterSize(f, 1<<16)
	for i := 1; i <= size; i++ {
		fmt.Fprintf(w, "num: %d\n", i)
	}
	w.Flush()
}
