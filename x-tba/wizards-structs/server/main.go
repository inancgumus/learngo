// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/add", onlyPost(add))
	http.HandleFunc("/list", list)
	panic(http.ListenAndServe(":8000", nil))
}

func add(w http.ResponseWriter, r *http.Request) {
	var wiz wizard
	if err := json.NewDecoder(r.Body).Decode(&wiz); err != nil {
		panic(err)
	}

	db.add(wiz)
	w.WriteHeader(http.StatusOK)
}

func list(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, db.list())
}

func onlyPost(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			next(w, r)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("NOT OK"))
	}
}
