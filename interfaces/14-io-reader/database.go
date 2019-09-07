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
	list  *list
	types map[string]item // the registry of the types
}

// #v1
// func (db *database) load(path string) error {
// 	data, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		return err
// 	}
// 	return json.Unmarshal(data, db)
// }

// #v2
// func (db *database) load(r io.Reader) error {
//  data, err := ioutil.ReadAll(r)
// 	if err != nil {
// 		return err
// 	}
// 	return json.Unmarshal(data, db)
// }

// #v3
// TODO: use decoder

// #v1
// func (db *database) save(path string) error {
// 	data, err := json.MarshalIndent(db, "", "\t")
// 	if err != nil {
// 		return err
// 	}
// 	return ioutil.WriteFile(path, data, 0644)
// }

// #v2
// func (db *database) save(w io.Writer) error {
// 	data, err := json.MarshalIndent(db, "", "\t")
// 	if err != nil {
// 		return err
// 	}
//
//  return io.Copy(w, bytes.NewReader(data))
// }

// #v3
// TODO: use encoder

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

func (db *database) newItem(data []byte, typ string) (item, error) {
	it, ok := db.types[typ]
	if !ok {
		return nil, fmt.Errorf("newItem: type (%q) does not exist", typ)
	}

	// get the dynamic type inside "it" like *book, *game, etc..
	t := reflect.TypeOf(it)

	// get the pointer type's element type like book, game, etc...
	e := t.Elem()

	// create a new pointer value of the type like new(book), new(game), etc...
	v := reflect.New(e)

	// get the interface value and cast it to the item type
	it = v.Interface().(item)

	return it, json.Unmarshal(data, &it)
}

func (db *database) register(typ string, it item) {
	if db.types == nil {
		db.types = make(map[string]item)
	}
	db.types[typ] = it
}
