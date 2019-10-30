// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "sync"

type wizard struct {
	Name     string `json:"name"`
	Lastname string `json:"last_name"`
	Nick     string `json:"nick"`
}

type storage struct {
	sync.RWMutex
	wizards []wizard
}

func (db *storage) add(w wizard) {
	db.Lock()
	defer db.Unlock()
	db.wizards = append(db.wizards, w)
}

func (db *storage) list() []wizard {
	db.RLock()
	defer db.RUnlock()
	return db.wizards
}

var db *storage

func init() {
	db = new(storage)
}
