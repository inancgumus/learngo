// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var start, stop int

	if args := os.Args[1:]; len(args) == 2 {
		start, _ = strconv.Atoi(args[0])
		stop, _ = strconv.Atoi(args[1])
	}

	if start == 0 || stop == 0 {
		start, stop = 'A', 'Z'
	}

	fmt.Printf("%-10s %-10s %-10s %-12s\n%s\n",
		"literal", "dec", "hex", "encoded",
		strings.Repeat("-", 45))

	for n := start; n <= stop; n++ {
		fmt.Printf("%-10c %-10[1]d %-10[1]x % -12x\n", n, string(n))
	}
}

/*
EXAMPLE UNICODE BLOCKS

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
latin1Start    = '\u0080'      ->  161
latin1Stop     = '\u00ff'      ->  255


3 bytes
------------------------------------------------------------
dingbatStart   = '\u2700'      ->  9984
dingbatStop    = '\u27bf'      ->  10175


4 bytes
------------------------------------------------------------
emojiStart     = '\U0001f600'  ->  128512
emojiStop      = '\U0001f64f'  ->  128591
*/
