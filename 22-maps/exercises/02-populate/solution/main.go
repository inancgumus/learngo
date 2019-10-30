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
	phones := map[string]string{
		"bowen": "202-555-0179",
		"dulin": "03.37.77.63.06",
		"greco": "03489940240",
	}

	products := map[int]bool{
		617841573: true,
		879401371: false,
		576872813: true,
	}

	multiPhones := map[string][]string{
		"bowen": {"202-555-0179"},
		"dulin": {"03.37.77.63.06", "03.37.70.50.05", "02.20.40.10.04"},
		"greco": {"03489940240", "03489900120"},
	}

	basket := map[int]map[int]int{
		100: {617841573: 4, 576872813: 2},
		101: {576872813: 5, 657473833: 20},
		102: {},
	}

	// Print dulin's phone number.
	who, phone := "dulin", "N/A"
	if v, ok := phones[who]; ok {
		phone = v
	}
	fmt.Printf("%s's phone number: %s\n", who, phone)

	// Is Product ID 879401371 available?
	id, status := 879401371, "available"
	if !products[id] {
		status = "not " + status
	}
	fmt.Printf("Product ID #%d is %s\n", id, status)

	// What is Greco's second phone number?
	who, phone = "greco", "N/A"
	if phones := multiPhones[who]; len(phones) >= 2 {
		phone = phones[1]
	}
	fmt.Printf("%s's 2nd phone number: %s\n", who, phone)

	// How many of 576872813 the customer 101 is going to buy?
	cid, pid := 101, 576872813
	fmt.Printf("Customer #%d is going to buy %d from Product ID #%d.\n", cid, basket[cid][pid], pid)
}
