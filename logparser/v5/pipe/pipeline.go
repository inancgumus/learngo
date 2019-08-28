// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package pipe

import (
	"fmt"
	"os"
)

// Pipeline takes records from a source, transforms, and sends them to a destionation.
type Pipeline struct {
	src   Iterator
	trans []Transform
	dst   Digester
}

// New creates a new pipeline.
func New(src Iterator, dst Digester, t ...Transform) *Pipeline {
	return &Pipeline{
		src:   &logCount{Iterator: src},
		dst:   dst,
		trans: t,
	}
}

// Run the pipeline.
func (p *Pipeline) Run() error {
	defer func() {
		n := p.src.(*logCount).count()
		fmt.Fprintf(os.Stderr, "%d records processed.\n", n)
	}()

	last := p.src

	for _, t := range p.trans {
		if err := t.Digest(last); err != nil {
			return err
		}
		last = t
	}

	return p.dst.Digest(last)
}
