// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	dec := make(decoder)
	dec["book"] = book{}
	dec["game"] = game{}
	dec["puzzle"] = puzzle{}
	dec["toy"] = toy{}

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	store, err := dec.decode(data)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print(store)

	// store.discount(.5)
	// fmt.Print(store)

	// ----
	// data, err := save(store)
	// fmt.Println(string(data), err)
}
