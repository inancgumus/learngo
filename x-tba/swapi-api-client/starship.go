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
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Starship represents a Star Wars space ship
type Starship struct {
	// All fields should be exported to be decoded
	Name    string      `json:"name"`
	Credits json.Number `json:"cost_in_credits"`
	Rating  float64     `json:"hyperdrive_rating,string"`
}

func (ship Starship) String() string {
	var buf strings.Builder

	fmt.Fprintf(&buf, "Ship Name: %s\n", ship.Name)
	fmt.Fprintf(&buf, "Price    : %s\n", ship.Credits)

	return buf.String()
}

// TODO: return error instead of quit()
func fetchStarships(ctx context.Context, shipUrls []string, ships []Starship) error {
	ships = ships[:0]

	re := regexp.MustCompile(`/([0-9]+)/$`)
	for _, url := range shipUrls {
		ids := re.FindAllStringSubmatch(url, -1)
		if ids == nil {
			continue
		}

		sid := ids[0][1]
		id, _ := strconv.Atoi(sid)

		// TODO: goroutine
		ship, err := fetchStarship(ctx, id)
		if err != nil {
			return err
		}

		ships = append(ships, ship)
	}
	return nil
}

// fetchStarship returns Starship info by its id
func fetchStarship(ctx context.Context, id int) (ship Starship, err error) {
	c, err := request(ctx, fmt.Sprintf(swapi+"starships/%d/", id))
	if err != nil {
		return ship, err
	}

	// -> If your response body is small enough,
	//    just read it all into memory using ioutil.ReadAll
	//    and use json.Unmarshal.
	//
	// -> Do not use json.Decoder if you are not dealing
	//    with JSON streaming.

	return ship, json.Unmarshal(c, &ship)
}
