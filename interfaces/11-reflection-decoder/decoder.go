// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "encoding/json"

type decoderSchema struct {
	Category string

	// json.RawMessage is a special type of the json package.
	// It prevents the json package functions encode/decode
	// the specific parts of the json data.
	Item json.RawMessage
}

func decode(data []byte) (list, error) {
	var schemas []decoderSchema

	if err := json.Unmarshal(data, &schemas); err != nil {
		return nil, err
	}

	return toStore(schemas)
}

func toStore(schemas []decoderSchema) (store list, err error) {
	for _, s := range schemas {
		p := newProduct(s.Category)

		if json.Unmarshal(s.Item, &p); err != nil {
			return nil, err
		}

		store = append(store, p)
	}
	return store, nil
}

func newProduct(category string) item {
	// for each product you need to open this code
	// and add a new case.
	//
	// dependency injection should be the responsibility
	// of the user of this function.
	switch category {
	case "book":
		// var b book
		// return &b
		return new(book)
	case "game":
		return new(game)
	case "puzzle":
		return new(puzzle)
	case "toy":
		return new(toy)
	}
	return nil
}

// func (d decoder) newProduct(category string) item {
// 	t := reflect.TypeOf(d[category])
// 	v := reflect.New(t)
// 	return v.Interface().(item)
// }

// func toProduct(kind string, data []byte) (item, error) {
// 	p := newProduct(kind)

// 	err := json.Unmarshal(data, p)

// 	return p, err
// }
