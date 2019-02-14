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
	"strconv"
	"strings"
)

const (
	asciiStart = 0x41 // 65
	asciiStop  = 0x5a // 90

	cols = 1
)

func main() {
	// DETERMINE START - STOP POSITIONS
	var start, stop int

	if args := os.Args[1:]; len(args) == 2 {
		start, _ = strconv.Atoi(args[0])
		stop, _ = strconv.Atoi(args[1])
	}

	if start == 0 || stop == 0 {
		start, stop = asciiStart, asciiStop
	}

	// PRINT HEADER
	for i := 0; i < cols; i++ {
		fmt.Printf("%-10s %-12s %-12s %-14s",
			"literal", "decimal", "codepoint", "bytes")
	}
	fmt.Print("\n", strings.Repeat("-", 50*cols), "\n")

	// PRINT TABLE
	for n, l := start, 0; n <= stop; n++ {
		// draw the line
		fmt.Printf("%-10q %-12d %-12U % -14x", n, n, n, string(n))

		// go to next line if columns are consumed
		if l++; l%cols == 0 {
			fmt.Println()
			continue
		}
	}
	fmt.Println()
}

/*
EXAMPLE BLOCKS

1 byte
------------------------------------------------------------
asciiStart     = '\u0001'      ->  32
asciiStop      = '\u007f'      ->  127

upperCaseStart = '\u0041'      ->  65
upperCaseStop  = '\u005a'      ->  90

lowerCaseStart = '\u0061'      ->  97
lowerCaseStop  = '\u007a'      ->  122


2 bytes
------------------------------------------------------------
latin1Start    = '\u0080'      ->  128
latin1Stop     = '\u00ff'      ->  255


3 bytes
------------------------------------------------------------
dingbatStart   = '\u2700'      ->  9984
dingbatStop    = '\u27bf'      ->  10175


4 bytes
------------------------------------------------------------
emojiStart     = '\U0001f600'  ->  128512
emojiStop      = '\U0001f64f'  ->  128591
transportStart = '\U0001F680'  ->  128640
transportStop  = '\U0001f6ff'  ->  128767

BIG THANK YOU! -> https://unicode-table.com/
*/
