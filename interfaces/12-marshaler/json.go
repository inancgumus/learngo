// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"encoding/json"
)

func encode() ([]byte, error) {
	l := list{
		{Title: "moby dick", Price: 10, Released: toTimestamp(118281600)},
		{Title: "odyssey", Price: 15, Released: toTimestamp("733622400")},
		{Title: "hobbit", Price: 25},
	}

	return json.MarshalIndent(l, "", "\t")
}

func decode(data []byte) (l list, err error) {
	if err := json.Unmarshal(data, &l); err != nil {
		return nil, err
	}
	return l, nil
}
