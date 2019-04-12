// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	// create a new struct type
	type person struct {
		name, lastname string
		age            int
	}

	picasso := person{"Pablo", "Picasso", 95}
	freud := person{name: "Sigmund", lastname: "Freud", age: 83}

	// picasso.age++
	// picasso.name = "Master"

	master := picasso
	// var master person
	// master.name = picasso.name
	// master.lastname = picasso.lastname
	// master.age = picasso.age

	master = freud
	master.name = "Master"
	master.age++

	for _, p := range []person{picasso, freud, master} {
		if p == picasso {
			// if p.name == picasso.name &&
			// 	p.lastname == picasso.lastname &&
			// 	p.age == picasso.age {

			// p is a copy
			p.name = "Raphael"

			fmt.Printf("%s's age is %d\n", p.name, p.age)
		}
	}

	// Picasso is still the same
	fmt.Printf("\n%s's age is %d\n", picasso.name, picasso.age)

	// type team struct{ scores []int }
	// _ = team{} == team{}
}
