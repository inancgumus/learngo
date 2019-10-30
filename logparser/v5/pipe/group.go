// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package pipe

import (
	"sort"
)

// GroupFunc represents a grouping func that returns a grouping key.
// The type alias frees us from binding to a named type.
type GroupFunc = func(Record) (key string)

// Group records by a key.
type Group struct {
	sum  map[string]record // metrics per group key
	keys []string          // unique group keys
	key  GroupFunc
}

// GroupBy returns a new Group.
// It takes a group func that returns a group key.
// The returned group will group the record using the key.
func GroupBy(key GroupFunc) *Group {
	return &Group{
		sum: make(map[string]record),
		key: key,
	}
}

// Consume the records for grouping.
func (g *Group) Consume(records Iterator) error {
	group := func(r Record) error {
		k := g.key(r)

		if _, ok := g.sum[k]; !ok {
			g.keys = append(g.keys, k)
		}

		g.sum[k] = r.sum(g.sum[k])

		return nil
	}

	return records.Each(group)
}

// Each sends the grouped and sorted records to upstream.
func (g *Group) Each(yield func(Record) error) error {
	sort.Strings(g.keys)

	for _, k := range g.keys {
		err := yield(Record{g.sum[k]})

		if err != nil {
			return err
		}
	}

	return nil
}
