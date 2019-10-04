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

type timestamp struct {
	time.Time
}

// timestamp knows how to decode itself from json.
//
// UnmarshalJSON is an implementation of the json.Unmarshaler interface.
// json.Unmarshal and json.Decode call this method.
func (ts *timestamp) UnmarshalJSON(data []byte) error {
	*ts = toTimestamp(string(data))
	return nil
}

// timestamp knows how to encode itself to json.
//
// MarshalJSON is an implementation of the json.Marshaler interface.
// json.Marshal and json.Encode call this method.
func (ts timestamp) MarshalJSON() (out []byte, err error) {
	return strconv.AppendInt(out, ts.Unix(), 10), nil
}

func (ts timestamp) String() string {
	if ts.IsZero() {
		return "unknown"
	}

	// Mon Jan 2 15:04:05 -0700 MST 2006
	const layout = "2006/01"
	return ts.Format(layout)
}

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
