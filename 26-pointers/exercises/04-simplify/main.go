// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Simplify the code
// HINT    : Remove the unnecessary pointer usages
// ---------------------------------------------------------

package main

import "fmt"

func main() {
	var schools []map[int]string

	schools = append(schools, make(map[int]string))
	load(&schools[0], &([]string{"batman", "superman"}))

	schools = append(schools, make(map[int]string))
	load(&schools[1], &([]string{"spiderman", "wonder woman"}))

	fmt.Println(schools[0])
	fmt.Println(schools[1])
}

func load(m *map[int]string, students *[]string) {
	for i, name := range *students {
		(*m)[i+1] = name
	}
}
