// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"strconv"
	"time"
)

type book struct {
	title     string
	price     money
	published interface{}
}

func (b book) print() {
	p := format(b.published)
	fmt.Printf("%-15s: %s - (%v)\n", b.title, b.price.string(), p)
}

func format(v interface{}) string {
	// book{title: "hobbit", price: 25},
	if v == nil {
		return "unknown"
	}

	var t int

	// book{title: "moby dick", price: 10, published: 118281600},
	if v, ok := v.(int); ok {
		t = v
	}

	// book{title: "odyssey", price: 15, published: "733622400"},
	if v, ok := v.(string); ok {
		t, _ = strconv.Atoi(v)
	}

	u := time.Unix(int64(t), 0)
	return u.String()
}
