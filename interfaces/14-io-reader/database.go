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
	"io"
	"io/ioutil"
	"reflect"
)

type database struct {
	list  *list
	types map[string]item // the registry of the types
}

// load the list by decoding the data from any data source.
func (db *database) load(r io.Reader) error {
	// ReadAll reads from `r` entirely into memory.
	// `data` contains the entire data in memory.
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, db)
}

func (db *database) MarshalJSON() ([]byte, error) {
	type encodable struct {
		Type string
		Item item
	}

	var e []encodable

	for _, it := range *db.list {
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

	t := reflect.TypeOf(it)
	e := t.Elem()
	v := reflect.New(e)
	it = v.Interface().(item)

	return it, json.Unmarshal(data, &it)
}

func (db *database) register(typ string, it item) {
	if db.types == nil {
		db.types = make(map[string]item)
	}
	db.types[typ] = it
}
