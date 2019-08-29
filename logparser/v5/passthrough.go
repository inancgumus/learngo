// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"github.com/inancgumus/learngo/logparser/v5/pipe"
)

type passThrough struct {
	pipe.Iterator
}

func (t *passThrough) Consume(results pipe.Iterator) error {
	t.Iterator = results
	return nil
}

func (t *passThrough) Each(yield func(pipe.Record) error) error {
	pass := func(r pipe.Record) error {
		// fmt.Println(r.Fields())
		// fmt.Println(r.Int("visits"))
		return yield(r)
	}

	return t.Iterator.Each(pass)
}
