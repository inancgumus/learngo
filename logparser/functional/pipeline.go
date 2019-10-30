// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "os"

type (
	processFn func(r result)
	inputFn   func(processFn) error
	outputFn  func([]result) error
	filterFn  func(result) (include bool)
	groupFn   func(result) (key string)
)

type pipeline struct {
	read   inputFn
	write  outputFn
	filter filterFn
	group  groupFn
}

func (p *pipeline) filterBy(f filterFn) *pipeline { p.filter = f; return p }
func (p *pipeline) groupBy(f groupFn) *pipeline   { p.group = f; return p }
func (p *pipeline) from(f inputFn) *pipeline      { p.read = f; return p }
func (p *pipeline) to(f outputFn) *pipeline       { p.write = f; return p }

func (p *pipeline) defaults() {
	if p.filter == nil {
		p.filter = noopFilter
	}

	if p.group == nil {
		p.group = domainGrouper
	}

	if p.read == nil {
		p.read = textReader(os.Stdin)
	}

	if p.write == nil {
		p.write = textWriter(os.Stdout)
	}
}

func (p *pipeline) start() error {
	p.defaults()

	// retrieve and process the lines
	sum := make(map[string]result)

	process := func(r result) {
		if !p.filter(r) {
			return
		}

		k := p.group(r)
		sum[k] = r.add(sum[k])
	}

	// return err from input reader
	if err := p.read(process); err != nil {
		return err
	}

	// prepare the results for outputting
	var out []result
	for _, res := range sum {
		out = append(out, res)
	}

	// return err from output reader
	return p.write(out)
}
