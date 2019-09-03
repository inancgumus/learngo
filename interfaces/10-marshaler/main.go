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
	store := list{
		&book{product{"moby dick", 10}, toTimestamp(118281600)},
		&book{product{"odyssey", 15}, toTimestamp("733622400")},
		&book{product{"hobbit", 25}, unknown},
		&puzzle{product{"rubik's cube", 5}},
		&game{product{"minecraft", 20}},
		&game{product{"tetris", 5}},
		&toy{product{"yoda", 150}},
	}

	out, err := json.MarshalIndent(store, "", "\t")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(out))

	// store.discount(.5)
	// fmt.Print(store)

	// ----

	// var ts timestamp
	// json.Unmarshal([]byte(`118281600`), &ts)
	// fmt.Println(ts)

	// json.Unmarshal([]byte(`"18281600"`), &ts)
	// fmt.Println(ts)

	// json.Unmarshal([]byte(`"incorrect"`), &ts)
	// fmt.Println(ts)
}
