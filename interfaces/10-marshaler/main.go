// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	// store := list{
	// 	&book{product{"moby dick", 10}, 118281600},
	// 	&book{product{"odyssey", 15}, "733622400"},
	// 	&book{product{"hobbit", 25}, nil},
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

	// store.discount(.5)
	// fmt.Print(store)

	encode()
}

func encode() {

	books := []*book{
		{product{"moby dick", 10}, toTime(118281600)},
		{product{"odyssey", 15}, toTime("733622400")},
		{product{"hobbit", 25}, zeroTimestamp},
	}

	for _, b := range books {
		fmt.Println(b)
	}

	out, err := json.MarshalIndent(books, "", "\t")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(out))
}

func decode() {
	// books = nil
	out := []byte(`[
	{
		"Title": "moby dick",
		"Price": 10,
		"Published": "118281600"
	},
	{
		"Title": "odyssey",
		"Price": 15,
		"Published": 733622400
	},
	{
		"Title": "hobbit",
		"Price": 25,
		"Published": -62135596800
	},
	{
		"Title": "poka",
		"Price": 25,
		"Published": "1983/03"
	},
	{
		"Title": "fuka",
		"Price": 25
	}	
]`)

	var books []*book

	err := json.Unmarshal(out, &books)
	if err != nil {
		log.Fatalln(err)
	}

	for _, b := range books {
		fmt.Println(b)
	}
}
