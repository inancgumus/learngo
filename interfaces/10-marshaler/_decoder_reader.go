// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"reflect"
)

type decoder map[string]interface{}

type schema struct {
	Category string
	Item     json.RawMessage
}

func (d decoder) fromFile(path string) (list, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return d.readFile(f)
}

func (d decoder) readFile(f *os.File) (list, error) {
	return d.read(f)
}

func (d decoder) read(r io.Reader) (list, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return d.decode(data)
}

func (d decoder) decode(data []byte) (list, error) {
	var schemas []schema

	if err := json.Unmarshal(data, &schemas); err != nil {
		return nil, err
	}

	return d.products(schemas)
}

func (d decoder) products(ss []schema) (list, error) {
	var store list

	for _, s := range ss {
		prod, err := d.product(s.Category, s.Item)

		if err != nil {
			return nil, err
		}

		store = append(store, prod)
	}

	return store, nil
}

func (d decoder) product(kind string, data []byte) (item, error) {
	p := d.newProduct(kind)

	err := json.Unmarshal(data, p)

	return p, err
}

func (d decoder) newProduct(kind string) item {
	t := reflect.TypeOf(d[kind])
	v := reflect.New(t)
	return v.Interface().(item)
}
