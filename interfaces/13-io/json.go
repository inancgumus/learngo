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
)

func encode() ([]byte, error) {
	l := list{
		{Title: "moby dick", Price: 10, Released: toTimestamp(118281600)},
		{Title: "odyssey", Price: 15, Released: toTimestamp("733622400")},
		{Title: "hobbit", Price: 25},
	}

	return json.MarshalIndent(l, "", "\t")
}

func decode(data []byte) (l list, _ error) {
	if err := json.Unmarshal(data, &l); err != nil {
		return nil, err
	}
	return
}

func decodeFile(path string) (list, error) {
	// ReadAll reads entire bytes from a file to memory.
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return decode(data)
}

func decodeFileObject(f *os.File) (list, error) {
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return decode(data)
}

func decodeReader(r io.Reader) (l list, err error) {
	// data, err := ioutil.ReadAll(r)
	// if err != nil {
	// 	return nil, err
	// }
	// return decode(data)

	return l, json.NewDecoder(r).Decode(&l)
}
