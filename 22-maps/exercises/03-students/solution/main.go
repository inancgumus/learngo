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
	"sort"
)

func main() {
	houses := map[string][]string{
		"gryffindor": {"weasley", "hagrid", "dumbledore", "lupin"},
		"hufflepuf":  {"wenlock", "scamander", "helga", "diggory", "bobo"},
		"ravenclaw":  {"flitwick", "bagnold", "wildsmith", "montmorency"},
		"slytherin":  {"horace", "nigellus", "higgs", "bobo", "scorpius"},
		"bobo":       {"wizardry", "unwanted"},
	}

	// remove the "bobo" house
	delete(houses, "bobo")

	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Please type a Hogwarts house name.")
		return
	}

	house, students := args[0], houses[args[0]]
	if students == nil {
		fmt.Printf("Sorry. I don't know anything about %q.\n", house)
		return
	}

	// only sort the clone
	clone := append([]string(nil), students...)
	sort.Strings(clone)

	fmt.Printf("~~~ %s students ~~~\n\n", house)
	for _, student := range clone {
		fmt.Printf("\t+ %s\n", student)
	}
}
