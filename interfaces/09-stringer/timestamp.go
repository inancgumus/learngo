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

// Mon Jan 2 15:04:05 -0700 MST 2006
const layout = "2006/01"

// unknown is the zero value of a timestamp.
var unknown = timestamp(time.Time{})

// Timestamp prints timestamps, it's a stringer.
// Timestamp is useful even if it's zero.
type timestamp time.Time

// String makes the timestamp a stringer.
func (ts timestamp) String() string {
	t := time.Time(ts)

	if t.IsZero() {
		return "unknown"
	}

	return t.Format(layout)
}

// toTimestamp was "book.format" before.
// Now it returns a timestamp value depending on the type of `v`.
func toTimestamp(v interface{}) timestamp {
	var t int

	switch v := v.(type) {
	case int:
		t = v
	case string:
		t, _ = strconv.Atoi(v)
	default:
		return unknown
	}

	return timestamp(time.Unix(int64(t), 0))
}
