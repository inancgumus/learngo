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
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Give me the magnitude of the earthquake.")
		return
	}

	richter, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		fmt.Println("I couldn't get that, sorry.")
		return
	}

	var desc string

	switch r := richter; {
	case r >= 10:
		desc = "massive"
	case r >= 8:
		desc = "great"
	case r >= 7:
		desc = "major"
	case r >= 6:
		desc = "strong"
	case r >= 5:
		desc = "moderate"
	case r >= 4:
		desc = "light"
	case r >= 3:
		desc = "minor"
	case r >= 2:
		desc = "very minor"
	default:
		desc = "micro"
	}

	fmt.Printf("%.2f is %s\n", richter, desc)
}
