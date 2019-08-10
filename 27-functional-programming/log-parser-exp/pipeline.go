package main

import (
	"fmt"
	"os"
	"time"
)

type (
	inputFn  func() ([]result, error)
	outputFn func([]result) error
	filterFn func(result) (include bool)
	groupFn  func(result) (key string)
)

type pipeline struct {
	input    inputFn
	filter   filterFn
	groupKey groupFn
	output   outputFn
}

func newPipeline() *pipeline {
	return &pipeline{
		filter:   noopFilter,
		groupKey: noopGrouper,
		input:    textReader(os.Stdin),
		output:   textWriter(os.Stdout),
	}
}

func (p *pipeline) from(fn inputFn) *pipeline      { p.input = fn; return p }
func (p *pipeline) to(fn outputFn) *pipeline       { p.output = fn; return p }
func (p *pipeline) filterBy(fn filterFn) *pipeline { p.filter = fn; return p }
func (p *pipeline) groupBy(fn groupFn) *pipeline   { p.groupKey = fn; return p }

func (p *pipeline) start() ([]result, error) {
	res, err := p.input()
	if err != nil {
		return nil, err
	}

	gres := make(map[string]result)

	for _, r := range res {
		if !p.filter(r) {
			continue
		}

		k := p.groupKey(r)
		gres[k] = r.add(gres[k])
	}

	var out []result
	for _, v := range gres {
		out = append(out, v)
	}

	err = p.output(out)

	return out, err
}

// TODO: remove me
func measure(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
