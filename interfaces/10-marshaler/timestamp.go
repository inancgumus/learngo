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

// unknown is the zero value of a timestamp.
var unknown = timestamp(time.Time{})

// timestamp prints timestamps, it's a stringer.
// It is useful even if it's zero.
type timestamp time.Time

// MarshalJSON is an implementation of the json.Marshaler interface.
// json.Marshal and json.Encode call this method.
func (ts timestamp) MarshalJSON() (out []byte, err error) {
	u := time.Time(ts).Unix()
	return strconv.AppendInt(out, u, 10), nil
}

// UnmarshalJSON is an implementation of the json.Unmarshaler interface.
// json.Unmarshal and json.Decode call this method.
func (ts *timestamp) UnmarshalJSON(data []byte) error {
	s := string(data)

	// Let the ParseInt parse quoted strings.
	us, err := strconv.Unquote(s)
	if err == nil {
		s = us
	}

	// Always overwrite the timestamp when decoding.
	*ts = unknown

	// Handle the numeric case.
	if n, err := strconv.ParseInt(s, 10, 64); err == nil {
		*ts = timestamp(time.Unix(n, 0))
	}

	return nil
}

func (ts timestamp) String() string {
	t := time.Time(ts)

	if t.IsZero() {
		return "unknown"
	}

	// Mon Jan 2 15:04:05 -0700 MST 2006
	const layout = "2006/01"

	return t.Format(layout)
}

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
