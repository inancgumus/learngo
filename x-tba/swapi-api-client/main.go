// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"
)

// const
// nil, string, int, float64, bool, comparison
// variables
//   multiple short variables
//   assignment?
// if
// error handling
// functions
//   returns
//   defer
// struct
//   encoding/json
// pointers
// concurrency
//   select
//   chan receive
// fmt
//   Printf
//   Sprintf
//   Errorf
// net/http
//   Get
// context/Context

// TODO: convert fmt calls to log
// TODO: you can make the fetcher a library and main package the user
// TODO: you can generate an html for the ship details? template pkg.

const timeout = 10 * time.Second

func main() {
	args := os.Args[1:]
	quit("give me a film id", len(args) != 1)

	id, err := strconv.Atoi(args[0])
	quit("film id is incorrrect", err)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// TODO: print the ship details to a text file (or any io.Writer)

	film, err := fetchFilm(ctx, id)
	quit("Error occurred while fetching the film data", err)
	fmt.Println(film)

	// a channel also can be used to print as they come
	ships := make([]Starship, len(film.Starships))
	err = fetchStarships(ctx, film.Starships, ships)
	quit("Error occurred while fetching starships", err)

	fmt.Println("Ships used in the movie:")
	fmt.Println("------------------------")
	for _, ship := range ships {
		fmt.Println(ship)
	}
}

func quit(message string, cond interface{}) {
	var quit bool

	switch v := cond.(type) {
	case error:
		quit = true
		message += ": " + v.Error()
	case bool:
		quit = v
	}

	if quit {
		fmt.Fprintln(os.Stderr, message)
		os.Exit(1)
	}
}
