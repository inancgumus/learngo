package main

import "io"

type (
	parserFunc func(io.Reader) ([]result, error)
	filterFunc func(result) bool
	groupFunc  func(result) string
	outputFunc func(io.Writer, []result) error
)

type report struct {
	input     io.Reader
	output    io.Writer
	parser    parserFunc
	filterer  filterFunc
	grouper   groupFunc
	outputter outputFunc
}

func newReport() *report {
	return &report{
		// parser:   textParser,
		filterer: noopFilter,
	}
}

func (r *report) from(reader io.Reader) *report {
	r.input = reader
	return r
}

func (r *report) to(writer io.Writer) *report {
	r.output = writer
	return r
}

func (r *report) retrieveFrom(fn parserFunc) *report {
	r.parser = fn
	return r
}

func (r *report) filterBy(fn filterFunc) *report {
	r.filterer = fn
	return r
}

func (r *report) groupBy(fn groupFunc) *report {
	r.grouper = fn
	return r
}

func (r *report) writeTo(fn outputFunc) *report {
	r.outputter = fn
	return r
}

func (r *report) run() ([]result, error) {
	if r.parser == nil {
		panic("report retriever cannot be nil")
	}

	results, err := r.parser(r.input)
	if err != nil {
		return nil, err
	}

	// noop if filterer is nil
	results = filterBy(results, r.filterer)

	// grouper is more tricky
	// you don't want to create an unnecessary map
	if r.grouper != nil {
		results = groupBy(results, r.grouper)
	}

	// prefer: noop output
	if r.output != nil {
		if err := r.outputter(r.output, results); err != nil {
			return nil, err
		}
	}

	return results, nil
}
