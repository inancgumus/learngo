// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"os"

	"github.com/inancgumus/learngo/internal/magic"
)

func main() {
	files := []string{
		"pngs/cups-jpg.png",
		"pngs/forest-jpg.png",
		"pngs/golden.png",
		"pngs/work.png",
		"pngs/shakespeare-text.png",
		"pngs/empty.png",
	}

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("png or jpg?")
		return
	}

	list(args[0], files)

	// fmt.Println("catch me if you can!")
}

func list(format string, files []string) {
	valids, err := magic.Detect(format, files)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Correct Files:\n")
	for _, valid := range valids {
		fmt.Printf(" + %s\n", valid)
	}
}
