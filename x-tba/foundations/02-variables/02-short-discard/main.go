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
	"math"
	"os"
	"path"
	"strconv"
)

func main() {
	var radius, area float64

	radius, _ = strconv.ParseFloat(os.Args[1], 64)

	area = 4 * math.Pi * math.Pow(radius, 2)
	// area := 4 * math.Pi * math.Pow(radius, 2)

	fmt.Printf("radius: %g -> area: %.2f\n",
		radius, area)

	dir, _ := path.Split("secret/file.txt")
	fmt.Println(dir)

	color, color2 := "red", "blue"
	color, color2 = color2, color
	fmt.Println(color, color2)
}
