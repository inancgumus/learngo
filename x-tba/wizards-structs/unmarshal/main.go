// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

// Wizard is one of the greatest of people
type Wizard struct {
	// name won't be marshalled (should be exported)
	Name     string `json:name`
	Lastname string `json:"-"`
	Nick     string `json:"nick"`
}

func main() {
	file, err := ioutil.ReadFile("../marshal/wizards.json")
	if err != nil {
		panic(err)
	}

	wizards := make([]Wizard, 10)
	if json.Unmarshal(file, &wizards) != nil {
		panic(err)
	}

	fmt.Printf("%-15s %-15s\n%s",
		"Name", "Nick", strings.Repeat("=", 25))

	for _, w := range wizards {
		fmt.Printf("%-15s %-15s\n", w.Name, w.Nick)
	}
}
