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
)

// database is responsible for encoding and decoding a list.
type database struct {
	list *list
}

func (db *database) UnmarshalJSON(data []byte) error {
	// Provides a concrete data structure that we can use to decode.
	// It is a slice of structs.
	var decodables []struct {
		Type string          // The product type.
		Item json.RawMessage // The raw item data (from database.json)
	}

	// First, decode the "Type" part.
	// Leave the "Item" as raw json data.
	if err := json.Unmarshal(data, &decodables); err != nil {
		return err
	}

	// Decode the "Item" part for each json object.
	for _, d := range decodables {
		it, err := db.newItem(d.Item, d.Type)
		if err != nil {
			return err
		}

		*db.list = append(*db.list, it)
	}

	return nil
}

// newItem decodes and returns a product type wrapped in an item iface value.
func (*database) newItem(data []byte, itemType string) (it item, _ error) {
	switch itemType {
	default:
		return nil, fmt.Errorf("newItem: type (%q) does not exist", itemType)
	case "book":
		it = new(book)
	case "puzzle":
		it = new(puzzle)
	case "game":
		it = new(game)
	case "toy":
		it = new(toy)
	}
	return it, json.Unmarshal(data, &it)
}
