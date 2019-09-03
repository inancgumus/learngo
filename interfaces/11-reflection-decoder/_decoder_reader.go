// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"io"
	"io/ioutil"
	"os"
)

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
