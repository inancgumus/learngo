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
	// store := list{
	// 	&book{product{"moby dick", 10}, toTimestamp(118281600)},
	// 	&book{product{"odyssey", 15}, toTimestamp("733622400")},
	// 	&book{product{"hobbit", 25}, unknown},
	// 	&puzzle{product{"rubik's cube", 5}},
	// 	&game{product{"minecraft", 20}},
	// 	&game{product{"tetris", 5}},
	// 	&toy{product{"yoda", 150}},
	// }

	// out, err := json.MarshalIndent(store, "", "\t")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// fmt.Println(string(out))

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
