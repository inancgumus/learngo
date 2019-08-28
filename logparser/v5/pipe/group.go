// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package pipe

import (
	"sort"
)

// GroupFunc represents a grouping func that returns a grouping key.
type GroupFunc = func(Record) (key string)

// Group records by a key.
type Group struct {
	sum  map[string]Record // metrics per group key
	keys []string          // unique group keys
	key  GroupFunc
}

// GroupBy returns a new Group.
// It takes a group func that returns a group key.
// The returned group will group the record using the key.
func GroupBy(key GroupFunc) *Group {
	return &Group{
		sum: make(map[string]Record),
		key: key,
	}
}

// Consume records for grouping.
func (g *Group) Consume(records Iterator) error {
	return records.Each(func(r Record) error {
		k := g.key(r)

		if _, ok := g.sum[k]; !ok {
			g.keys = append(g.keys, k)
		}

		g.sum[k] = r.Sum(g.sum[k])

		return nil
	})
}

// Each sorts and yields the grouped records.
func (g *Group) Each(yield func(Record) error) error {
	sort.Strings(g.keys)

	for _, k := range g.keys {
		if err := yield(g.sum[k]); err != nil {
			return err
		}
	}

	return nil
}
