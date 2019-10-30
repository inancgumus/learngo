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
	var (
		myAge  byte
		herAge byte

		// this declares an array with two byte elements
		// its length      : 2
		// its element type: byte
		ages [2]byte

		// this declares an array with five string elements
		// its length      : 5
		// its element type: string
		tags [5]string

		// this array doesn't occupy any memory space (its length is zero)
		// its length      : 0
		// its element type: byte
		zero [0]byte

		// this array uses a constant expression
		// its length      : 3
		// its element type: byte
		agesExp [2 + 1]byte

		// uncomment the code below and observe the error
		//
		// wrongDeclaration [-1]byte
	)

	fmt.Printf("myAge             : %d\n", myAge)
	fmt.Printf("herAge            : %d\n", herAge)

	// Since arrays we've declared here don't have any elements,
	// Go automatically sets all their elements to their zero values
	// depending on their element type.

	// #v verb prints the array's length, element type and its elements
	fmt.Printf("ages              : %#v\n", ages)
	fmt.Printf("tags              : %#v\n", tags)
	fmt.Printf("zero              : %#v\n", zero)
	fmt.Printf("agesExp           : %#v\n", agesExp)

	// note:
	// ages and agesExp get printed 0x0 because they're byte arrays.
	// bytes are represented with hex values. 0x0 means 0.

	// =========================================================================
	// GETTING AND SETTING ARRAY ELEMENTS
	// =========================================================================

	// Note:
	//
	// Since, I've already declared the ages variable above,
	//   and, to show you the example below, I needed to create a new block.
	//
	// ages variable below is in a new block below. So, it's a new variable.
	//
	// I did so because I need to change the element type of the ages array
	//   to int (or, subtracting from a byte results in wraparound).
	{
		var ages [2]int

		fmt.Println()
		fmt.Printf("ages              : %#v\n", ages)
		fmt.Printf("ages's type       : %T\n", ages)

		fmt.Println("len(ages)         :", len(ages))
		fmt.Println("ages[0]           :", ages[0])
		fmt.Println("ages[1]           :", ages[1])
		fmt.Println("ages[len(ages)-1] :", ages[len(ages)-1])

		// WRONG:
		// fmt.Println("ages[-1]          :", ages[-1])
		// fmt.Println("ages[2]           :", ages[2])
		// fmt.Println("ages[len(ages)]   :", ages[len(ages)])

		ages[0] = 6
		ages[1] -= 3

		// WRONG:
		// ages[0] = "Can I?"

		fmt.Println("ages[0]           :", ages[0])
		fmt.Println("ages[1]           :", ages[1])

		ages[0] *= ages[1]
		fmt.Println("ages[0]           :", ages[0])
		fmt.Println("ages[1]           :", ages[1])
	}
}
