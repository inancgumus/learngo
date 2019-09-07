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
	"reflect"
)

type database struct {
	list *list
}

func (db *database) MarshalJSON() ([]byte, error) {
	type encodable struct {
		Type string
		Item item
	}

	var e []encodable

	for _, it := range *db.list {
		// TypeOf -> finds the dynamic type of "it"
		// Elem   -> returns the element type of the pointer
		// Name   -> returns the type name as string
		t := reflect.TypeOf(it).Elem().Name()
		e = append(e, encodable{t, it})
	}

	return json.Marshal(e)
}

func (db *database) UnmarshalJSON(data []byte) error {
	var decodables []struct {
		Type string
		Item json.RawMessage
	}

	if err := json.Unmarshal(data, &decodables); err != nil {
		return err
	}

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
