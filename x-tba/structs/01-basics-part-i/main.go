// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	/*
		USING VARIABLES
	*/

	// var (
	// 	name, lastname string
	// 	age            int
	// )
	//
	// name, lastname = "Pablo", "Picasso"
	// age = 95
	//
	// name2, lastname2 := "Sigmund", "Freud"
	// age2 := 83
	//
	// fmt.Println("Picasso:", name, lastname, age)
	// fmt.Println("Freud  :", name2, lastname2, age2)

	/*
		USING STRUCT VARIABLES
	*/

	// // declare a struct variable
	// var picasso struct {
	// 	name, lastname string
	// 	age            int
	// }

	// create a new struct type
	type person struct {
		name, lastname string
		age            int
	}

	// short-declare struct variable without field names
	// you cannot omit values when you initialize without field names
	picasso := person{
		name:     "Pablo",
		lastname: "Picasso",
		age:      95,
	}

	// short-declare struct variable with field names
	// omitted fields will be zeroed
	freud := person{
		name:     "Sigmund",
		lastname: "Freud",
		age:      83,
	}

	fmt.Printf("Picasso: %#v\n\n", picasso)
	fmt.Printf("Freud  : %#v\n\n", freud)
}
