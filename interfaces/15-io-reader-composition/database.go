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
	"reflect"
)

type database struct {
	list  *list
	types map[string]item
}

// load the list by decoding the data from any kind of data source.
func (db *database) load(r io.Reader) error {
	return json.NewDecoder(r).Decode(db)
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

// I put it here as a reference implementation of manually reading from a reader.
// func (db *database) loadX(r io.Reader) error {
// 	buf := make([]byte, 64)
//
// 	var (
// 		data []byte
// 		n    int
// 		err  error
// 	)
//
// 	for {
// 		n, err = r.Read(buf)
//
// 		if n > 0 {
// 			data = append(data, buf[:n]...)
//		}
//
// 		if err != nil {
// 			break
// 		}
// 	}
//
// 	if err != io.EOF {
// 		return err
// 	}
// 	return json.Unmarshal(data, db)
// }
