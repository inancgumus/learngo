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

	"github.com/inancgumus/learngo/09-go-type-system/05-defined-types/03-underlying-types/weights"
)

type (
	// Gram underlying type is int
	Gram int

	// Kilogram underlying type is int
	Kilogram Gram

	// Ton underlying type is int
	Ton Kilogram
)

func main() {
	var (
		salt   Gram     = 100
		apples Kilogram = 5
		truck  Ton      = 10
	)

	// types with different names cannot be used together
	// salt = apples
	// apples = truck

	// since their underlying type is int
	// they can be converted to any type
	//   with an underlying type of int
	salt = Gram(apples)
	apples = Kilogram(truck)
	truck = Ton(Kilogram(Gram(int(apples))))

	fmt.Printf("salt: %d, apples: %d, truck: %d\n",
		salt, apples, truck)

	// -----------------------------------------------------
	// TYPES FROM DIFFERENT PACKAGES
	// -----------------------------------------------------

	// Gram and weights.Gram are different types

	// You cannot do this
	// salt = weights.Gram(100)

	// But, you can convert it back to Gram
	// Since, their underlying type is int
	salt = Gram(weights.Gram(100))

	fmt.Printf("Type of weights.Gram: %T\n", weights.Gram(1))
	fmt.Printf("Type of main.Gram   : %T\n", Gram(1))

	// You can declare a new type
	// from a type of an external package
	type myGram weights.Gram

	var pepper myGram = 20
	pepper = myGram(salt)

	fmt.Printf("Type of pepper      : %T\n", pepper)
}
