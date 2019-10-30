// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

const data = `
[
        {
                "id": 1,
                "name": "god of war",
                "genre": "action adventure",
                "price": 50
        },
        {
                "id": 2,
                "name": "x-com 2",
                "genre": "strategy",
                "price": 40
        },
        {
                "id": 3,
                "name": "minecraft",
                "genre": "sandbox",
                "price": 20
        }
]`

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

func load() (games []game) {
	games = addGame(games, newGame(1, 50, "god of war", "action adventure"))
	games = addGame(games, newGame(2, 40, "x-com 2", "strategy"))
	games = addGame(games, newGame(3, 20, "minecraft", "sandbox"))
	return
}

func addGame(games []game, g game) []game {
	return append(games, g)
}

func newGame(id, price int, name, genre string) game {
	return game{
		item:  item{id: id, name: name, price: price},
		genre: genre,
	}
}

func indexByID(games []game) (byID map[int]game) {
	byID = make(map[int]game)
	for _, g := range games {
		byID[g.id] = g
	}
	return
}

func printGame(g game) {
	fmt.Printf("#%d: %-15q %-20s $%d\n",
		g.id, g.name, "("+g.genre+")", g.price)
}
