package main

import "os"

type (
	inputFunc  func() ([]result, error)
	outputFunc func([]result) error
)

type report struct {
	input  inputFunc
	filter filterFunc
	group  groupFunc
	output outputFunc
}

func newReport() *report {
	return &report{
		filter: noopFilter,
		input:  textReader(os.Stdin),
		output: textWriter(os.Stdout),
	}
}

func (r *report) from(fn inputFunc) *report {
	r.input = fn
	return r
}

func (r *report) to(fn outputFunc) *report {
	r.output = fn
	return r
}

func (r *report) filterBy(fn filterFunc) *report {
	r.filter = fn
	return r
}

func (r *report) groupBy(fn groupFunc) *report {
	r.group = fn
	return r
}

func (r *report) start() ([]result, error) {
	if r.input == nil {
		panic("report input cannot be nil")
	}

	results, err := r.input()
	if err != nil {
		return nil, err
	}

	// noop if filter is nil
	results = filterBy(results, r.filter)

	// group func is more tricky
	// you don't want to create an unnecessary map
	if r.group != nil {
		results = groupBy(results, r.group)
	}

	// TODO: prefer: noop writer
	if r.output != nil {
		if err := r.output(results); err != nil {
			return nil, err
		}
	}

	return results, nil
}
