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
	"time"
)

func main() {
	// single()
	// stacked()
	findTheMeaning()
}

func findTheMeaningNoDefer() {
	start := time.Now()
	fmt.Printf("%s starts...\n", "findTheMeaning")

	// do some heavy calculation...
	time.Sleep(time.Second * 2)

	fmt.Printf("%s took %v\n", "findTheMeaning", time.Since(start))
}

func findTheMeaning() {
	defer measure("findTheMeaning")()

	// do some heavy calculation
	time.Sleep(time.Second * 2)
}

func measure(name string) func() {
	start := time.Now()
	fmt.Printf("%s starts...\n", name)

	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func stacked() {
	for count := 1; count <= 5; count++ {
		defer fmt.Println(count)
	}

	fmt.Println("the stacked func returns")
}

func single() {
	var count int

	// defer func() {
	// 	fmt.Println(count)
	// }()
	defer fmt.Println(count)

	count++
	// fmt.Println(count)

	// the defer runs here
	// fmt.Println(count)
}
