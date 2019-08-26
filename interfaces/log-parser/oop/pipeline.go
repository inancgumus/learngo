// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"os"
	"strings"
)

type resultFn func(result)

type iterator interface {
	each(resultFn) error
}

type transformer interface {
	transform(result)
	iterator
}

type reporter interface {
	report(iterator) error
}

type pipeline struct {
	src  iterator
	dst  reporter
	tran transformer
}

func newPipeline(source iterator, r reporter, t transformer) *pipeline {
	return &pipeline{
		src:  source,
		dst:  r,
		tran: t,
	}
}

// fromFile generates a default report
func fromFile(path string) (err error) {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	var src iterator
	switch {
	case strings.HasSuffix(path, ".txt"):
		src = newTextLog(f)
	case strings.HasSuffix(path, ".json"):
		src = newJSONLog(f)
	}

	p := newPipeline(src, newTextReport(), newAnalysis())
	return p.run()
}

func (p *pipeline) run() error {
	if err := p.src.each(p.tran.transform); err != nil {
		return err
	}
	return p.dst.report(p.tran)
}
