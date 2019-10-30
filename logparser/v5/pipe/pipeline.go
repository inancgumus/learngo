// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package pipe

import (
	"fmt"
	"io"
	"os"
)

// Pipeline takes records from a source, transforms, and sends them to a destionation.
type Pipeline struct {
	src   Iterator
	trans []Transform
	dst   Consumer
}

// New creates a new pipeline.
func New(src Iterator, dst Consumer, t ...Transform) *Pipeline {
	return &Pipeline{
		src:   &logCount{Iterator: src},
		dst:   dst,
		trans: t,
	}
}

// Default creates a pipeline that reads from a text log and generates a text report.
func Default(r io.Reader, w io.Writer, t ...Transform) *Pipeline {
	return New(NewTextLog(r), NewTextReport(w), t...)
}

// Run the pipeline.
func (p *Pipeline) Run() error {
	defer func() {
		n := p.src.(*logCount).count()
		fmt.Fprintf(os.Stderr, "%d records processed.\n", n)
	}()

	last := p.src

	for _, t := range p.trans {
		if err := t.Consume(last); err != nil {
			return err
		}
		last = t
	}

	return p.dst.Consume(last)
}
