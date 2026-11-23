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
)

func main() {
	args := os.Args
	if len(args) != 2 && len(args) != 3 {
		fmt.Println("Tell me the magnitude of the earthquake in human terms.")
		return
	}

	var richter string

	desc := args[1]

	switch desc {
	case "micro":
		richter = "less than 2.0"
	case "minor":
		richter = "3 - 3.9"
	case "light":
		richter = "4 - 4.9"
	case "moderate":
		richter = "5 - 5.9"
	case "strong":
		richter = "6 - 6.9"
	case "major":
		richter = "7 - 7.9"
	case "great":
		richter = "8 - 9.9"
	case "massive":
		richter = "10+"
	default:
		richter = "unknown"
	}

	/*
		Since os.Args[1] will get only 'very' string, we need to
		concatenate os.Args[1] and os.Args[2] to form 'very minor'
	*/
	if len(args) == 3 && args[1]+" "+args[2] == "very minor" {
		richter = "2 - 2.9"
		desc = args[1] + " " + args[2]
	}

	fmt.Printf(
		"%s's richter scale is %s\n",
		desc, richter,
	)
}
