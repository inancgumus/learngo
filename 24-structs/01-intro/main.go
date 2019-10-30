// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

func main() {
	type Movie struct {
		Title  string
		Genre  string
		Rating int
	}

	type Rental struct {
		Address string
		Rooms   int
		Size    int
		Price   int
	}

	type Person struct {
		Name     string
		Lastname string
		Age      int
	}

	person1 := Person{Name: "Pablo", Lastname: "Picasso", Age: 91}
	person2 := Person{Name: "Sigmund", Lastname: "Freud", Age: 83}

	fmt.Printf("person1: %+v\n", person1)
	fmt.Printf("person2: %+v\n", person2)

	type VideoGame struct {
		Title     string
		Genre     string
		Published bool
	}

	pacman := VideoGame{
		Title:     "Pac-Man",
		Genre:     "Arcade Game",
		Published: true,
	}

	fmt.Printf("pacman: %+v\n", pacman)
}
