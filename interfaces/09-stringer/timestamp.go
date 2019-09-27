// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"strconv"
	"time"
)

// timestamp stores, formats and automatically prints a timestamp: it's a stringer.
type timestamp struct {
	// timestamp embeds a time, therefore it can be used as a time value.
	// there is no need to convert a time value to a timestamp value.
	time.Time
}

// String method makes the timestamp an fmt.stringer.
func (ts timestamp) String() string {
	if ts.IsZero() {
		return "unknown"
	}

	// Mon Jan 2 15:04:05 -0700 MST 2006
	const layout = "2006/01"
	return ts.Format(layout)
}

// toTimestamp returns a timestamp value depending on the type of `v`.
// toTimestamp was "book.format()" before.
func toTimestamp(v interface{}) timestamp {
	var t int

	switch v := v.(type) {
	case int:
		// book{title: "moby dick", price: 10, published: 118281600},
		t = v
	case string:
		// book{title: "odyssey", price: 15, published: "733622400"},
		t, _ = strconv.Atoi(v)
	default:
		// book{title: "hobbit", price: 25},
		return timestamp{}
	}

	return timestamp{
		Time: time.Unix(int64(t), 0),
	}
}
