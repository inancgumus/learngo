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
	"strconv"
	"time"
)

type book struct {
	product
	Published interface{}
}

func (b book) String() string {
	p := format(b.Published)
	return fmt.Sprintf("%s - (%v)", &b.product, p)
}

// MarshalJSON is an implementation of the json.Marshaler interface.
// json.Marshal and Decode call this method.
func (b *book) MarshalJSON() ([]byte, error) {
	// Calling Marshal creates an infinite loop because Marshal calls
	// MarshalJSON again and again.
	//
	// Declaring a clone type without MarshalJSON fixes the problem.
	type jbook book

	// book is convertable to jbook
	// jbook's underlying type is book.
	jb := (*jbook)(b)

	// Customize the formatting of the published field.
	// This will result in the published field of the resulting json.
	jb.Published = format(b.Published)

	return json.Marshal(jb)
}

func format(v interface{}) string {
	var t int

	switch v := v.(type) {
	case int:
		t = v
	case string:
		t, _ = strconv.Atoi(v)
	default:
		return "unknown"
	}

	// Mon Jan 2 15:04:05 -0700 MST 2006
	const layout = "2006/01"

	u := time.Unix(int64(t), 0)
	return u.Format(layout)
}
