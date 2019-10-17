// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"compress/gzip"
	"fmt"
	"net/http"
)

func main() {
	/*
		store := list{
			&book{product{"moby dick", 10}, toTimestamp(118281600)},
			&book{product{"odyssey", 15}, toTimestamp("733622400")},
			&book{product{"hobbit", 25}, unknown},
			&puzzle{product{"rubik's cube", 5}},
			&game{product{"minecraft", 20}},
			&game{product{"tetris", 5}},
			&toy{product{"yoda", 150}},
		}

		db := database{list: &store}

		out, err := json.MarshalIndent(&db, "", "\t")
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(string(out))

		// store.discount(.5)
		// fmt.Print(store)
	*/

	var store list

	db := database{list: &store}
	db.register("book", new(book))
	db.register("game", new(game))
	db.register("puzzle", new(puzzle))
	db.register("toy", new(toy))

	// load from a file
	// f, _ := os.Open("database.json")
	// db.load(f)
	// f.Close()

	// load from a string
	// const data = `[
	// { "Type": "book", "Item": { "Title": "1984", "Price": 8, "Published": -649641600 } },
	// { "Type": "game", "Item": { "Title": "paperboy", "Price": 20 } }]`
	// r := strings.NewReader(data)
	// db.load(r)

	// load from a web server
	// res, _ := http.Get("https://inancgumus.github.io/x/database.json")
	// db.load(res.Body)
	// res.Body.Close()

	// decompress the reader while reading from it
	// res, _ := http.Get("https://inancgumus.github.io/x/database.json.gz")
	// gr, _ := gzip.NewReader(res.Body)
	// db.load(gr)
	// gr.Close()
	// res.Body.Close()

	// decompress and count the number of compressed bytes read
	// res, _ := http.Get("https://inancgumus.github.io/x/database.json.gz")
	// c := newCounter(res.Body) // count the bytes read
	// gr, _ := gzip.NewReader(c)

	// db.load(gr)
	// fmt.Printf("%d bytes read.\n\n", c.total())

	// gr.Close()
	// res.Body.Close()

	// count the number of compressed bytes read from the web server then decompress
	res, _ := http.Get("https://inancgumus.github.io/x/database.json.gz")
	gr, _ := gzip.NewReader(res.Body)
	c := newCounter(gr)

	db.load(c)
	fmt.Printf("%d bytes read.\n\n", c.total())

	gr.Close()
	res.Body.Close()

	fmt.Print(store)
}
