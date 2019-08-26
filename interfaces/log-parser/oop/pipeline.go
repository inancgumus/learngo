// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"os"
	"strings"
)

type resultFn func(result)

type iterator interface{ each(resultFn) error }
type digester interface{ digest(iterator) error }

type transform interface {
	digester
	iterator
}

type pipeline struct {
	src   iterator
	trans []transform
	dst   digester
}

func (p *pipeline) run() error {
	defer func() {
		n := p.src.(*logCount).count()
		fmt.Printf("%d records processed.\n", n)
	}()

	last := p.src

	for _, t := range p.trans {
		if err := t.digest(last); err != nil {
			return err
		}
		last = t
	}

	return p.dst.digest(last)
}

func newPipeline(src iterator, dst digester, t ...transform) *pipeline {
	return &pipeline{
		src:   &logCount{iterator: src},
		dst:   dst,
		trans: t,
	}
}

// fromFile generates a default report
func fromFile(path string) (*pipeline, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var src iterator
	switch {
	case strings.HasSuffix(path, ".txt"):
		src = newTextLog(f)
	case strings.HasSuffix(path, ".jsonl"):
		src = newJSONLog(f)
	}

	return newPipeline(
		src,
		newTextReport(),
		groupBy(domainGrouper),
	), nil
}
