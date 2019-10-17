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
	out := encode()
	decode(out)
}

func decode(data []byte) {
	// data := []byte(`[
	//     { "Title": "moby dick", "Price": 5, "Published": 118281600 },
	//     { "Title": "odyssey", "Price": 7.5, "Published": 733622400 },
	//     { "Title": "hobbit", "Price": 12.5, "Published": -62135596800 }
	// ]`)

	var books []*book

	if err := json.Unmarshal(data, &books); err != nil {
		log.Fatalln(err)
	}

	for _, b := range books {
		fmt.Println(b)
	}
}

func encode() []byte {
	store := list{
		&book{product{"moby dick", 10}, toTimestamp(118281600)},
		&book{product{"odyssey", 15}, toTimestamp("733622400")},
		&book{product{"hobbit", 25}, toTimestamp(nil)},
		&puzzle{product{"rubik's cube", 5}},
		&game{product{"minecraft", 20}},
		&game{product{"tetris", 5}},
		&toy{product{"yoda", 150}},
	}

	store.discount(.5)
	// fmt.Print(store)

	out, err := json.MarshalIndent(store[:3], "", "\t")
	if err != nil {
		log.Fatalln(err)
	}

	return out
	// fmt.Println(string(out))
}
