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
	"strings"
)

type item interface {
	discount(float64)
	fmt.Stringer
}

// itemEncode adds an additional category field
// to the product types.
type itemEncode struct {
	// Category is the product category field.
	Category string

	// Item is an interface value.
	// Here, we are embedding an interface value.
	Item item
}

// itemDecode is used to decode the product types from json.
type itemDecode struct {
	// Category is the product category field.
	Category string

	// json.RawMessage is a special type of the json package.
	// It prevents the json package functions encode/decode
	// the specific parts of the json data.
	Item json.RawMessage
}

type list []item

func (l list) MarshalJSON() ([]byte, error) {
	var eitems []itemEncode

	for _, it := range l {
		// TypeOf -> finds the dynamic type of "it"
		// Elem   -> returns the element type of the pointer
		// Name   -> returns the type name as string
		cat := reflect.TypeOf(it).Elem().Name()
		eitems = append(eitems, itemEncode{cat, it})
	}

	return json.Marshal(eitems)
}

func (l *list) UnmarshalJSON(data []byte) error {
	var ditems []itemDecode

	if err := json.Unmarshal(data, &ditems); err != nil {
		return err
	}

	items, err := decode(ditems)
	if err != nil {
		return err
	}

	*l = items
	return nil
}

// decode and return items as a list.
func decode(data []itemDecode) (items list, err error) {
	for _, d := range data {
		item := newItem(d.Category)

		if json.Unmarshal(d.Item, &item); err != nil {
			return nil, err
		}

		items = append(items, item)
	}
	return items, nil
}

// newItem creates and returns an item depending on category.
func newItem(category string) item {
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

// type decoder map[string]interface{}
// func (d decoder) newItem(kind string) item {
// 	t := reflect.TypeOf(d[kind])
// 	v := reflect.New(t)
// 	return v.Interface().(item)
// }

func (l list) String() string {
	if len(l) == 0 {
		return "Sorry. We're waiting for delivery 🚚."
	}

	var str strings.Builder

	for _, it := range l {
		str.WriteString(it.String())
		str.WriteRune('\n')
	}

	return str.String()
}

func (l list) discount(ratio float64) {
	for _, it := range l {
		it.discount(ratio)
	}
}
