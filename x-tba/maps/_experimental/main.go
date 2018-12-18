// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

// why maps instead of keyed arrays etc?
// 120k elements!

package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("[emoji]")
		return
	}

	moods := [...]string{
		'😊': "a happy",
		'😟': "a worried",
		'😎': "an awesome",
		'😞': "a sad",
		'😩': "a crying",
	}

	emoji, _ := utf8.DecodeRune([]byte(args[0]))
	fmt.Printf("%s is %s emoji\n", args[0], moods[emoji])
}
