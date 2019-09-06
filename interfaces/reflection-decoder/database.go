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
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

type database struct {
	list  *list
	types map[string]item
}

func (db *database) save() error {
	data, err := json.MarshalIndent(db, "", "\t")
	if err != nil {
		log.Fatalln(err)
	}

	return ioutil.WriteFile("database.json", data, os.ModePerm)
}

func (db *database) load() error {
	data, err := ioutil.ReadFile("database.json")
	if err != nil {
		return err
	}

	if err = json.Unmarshal(data, db); err != nil {
		return err
	}

	return nil
}

func (db *database) MarshalJSON() ([]byte, error) {
	type encodable struct {
		Category string
		Item     item
	}

	var e []encodable

	for _, it := range *db.list {
		// TypeOf -> finds the dynamic type of "it"
		// Elem   -> returns the element type of the pointer
		// Name   -> returns the type name as string
		cat := reflect.TypeOf(it).Elem().Name()
		e = append(e, encodable{cat, it})
	}

	return json.Marshal(e)
}

func (db *database) UnmarshalJSON(data []byte) error {
	var decodables []struct {
		Category string
		Item     json.RawMessage
	}

	if err := json.Unmarshal(data, &decodables); err != nil {
		return err
	}

	for _, d := range decodables {
		it, err := db.newItem(d.Category)
		if err != nil {
			return err
		}

		if err := json.Unmarshal(d.Item, &it); err != nil {
			return err
		}

		*db.list = append(*db.list, it)
	}

	return nil
}

func (db *database) register(category string, t item) {
	if db.types == nil {
		db.types = make(map[string]item)
	}
	db.types[category] = t
}

func (db *database) newItem(category string) (item, error) {
	it, ok := db.types[category]
	if !ok {
		return nil, fmt.Errorf("register %q within database before decoding", category)
	}

	t := reflect.TypeOf(it).Elem()
	v := reflect.New(t)
	return v.Interface().(item), nil
}
