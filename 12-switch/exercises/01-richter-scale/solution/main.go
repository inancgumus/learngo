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
	case r < 2:
		desc = "micro"
	case r >= 2 && r < 3:
		desc = "very minor"
	case r >= 3 && r < 4:
		desc = "minor"
	case r >= 4 && r < 5:
		desc = "light"
	case r >= 5 && r < 6:
		desc = "moderate"
	case r >= 6 && r < 7:
		desc = "strong"
	case r >= 7 && r < 8:
		desc = "major"
	case r >= 8 && r < 10:
		desc = "great"
	default:
		desc = "massive"
	}

	fmt.Printf("%.2f is %s\n", richter, desc)
}
