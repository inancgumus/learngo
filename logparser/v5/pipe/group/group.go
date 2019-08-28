// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package group

import (
	"sort"

	"github.com/inancgumus/learngo/logparser/v5/pipe"
)

// Func represents a grouping func that returns a grouping key.
type Func func(pipe.Record) (key string)

// Group records by a key.
type Group struct {
	sum  map[string]pipe.Record // metrics per group key
	keys []string               // unique group keys
	key  Func
}

// By returns a new Group.
// It takes a group func that returns a group key.
// The returned group will group the record using the key.
func By(key Func) *Group {
	return &Group{
		sum: make(map[string]pipe.Record),
		key: key,
	}
}

// Consume records for grouping.
func (g *Group) Consume(records pipe.Iterator) error {
	return records.Each(func(r pipe.Record) {
		k := g.key(r)

		if _, ok := g.sum[k]; !ok {
			g.keys = append(g.keys, k)
		}

		if r, ok := r.(pipe.Summer); ok {
			g.sum[k] = r.Sum(g.sum[k])
		}
	})
}

// Each sorts and yields the grouped records.
func (g *Group) Each(yield func(pipe.Record)) error {
	sort.Strings(g.keys)

	for _, k := range g.keys {
		yield(g.sum[k])
	}

	return nil
}
