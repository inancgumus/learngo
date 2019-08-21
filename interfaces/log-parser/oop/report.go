// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"errors"
	"os"
	"strings"
)

type resultFn func(result)

type parser interface {
	parse(resultFn) error
}

type analyser interface {
	analyse(result)
}

type iterator interface {
	each(resultFn)
}

type summarizer interface {
	summarize(iterator) error
}

func report(p parser, a analyser, s summarizer) error {
	if err := p.parse(a.analyse); err != nil {
		return err
	}

	it, ok := a.(iterator)
	if !ok {
		return errors.New("cannot iterate on analyser")
	}

	return s.summarize(it)
}

// reportFromFile generates a default report
func reportFromFile(path string) (err error) {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	var p parser
	switch {
	case strings.HasSuffix(path, ".txt"):
		p = newTextParser(f)
	case strings.HasSuffix(path, ".json"):
		p = newJSONParser(f)
	}

	return report(p, newAnalysis(), newTextSummary())
}
