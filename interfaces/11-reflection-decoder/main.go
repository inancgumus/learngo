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
)

func main() {
	data, err := ioutil.ReadFile("database.json")
	if err != nil {
		log.Fatalln(err)
	}

	store, err := decode(data)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print(store)
	// for _, p := range store {
	// 	fmt.Printf("%#v\n", p)
	// }
}
