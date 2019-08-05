package main

import "io"

type (
	retrieveFunc func(io.Reader) ([]result, error)
	writeFunc    func(io.Writer, []result) error
)

type report struct {
	input    io.Reader
	output   io.Writer
	retrieve retrieveFunc
	filter   filterFunc
	group    groupFunc
	write    writeFunc
}

func newReport() *report {
	return &report{
		filter: noopFilter,
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

func (r *report) retrieveFrom(fn retrieveFunc) *report {
	r.retrieve = fn
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

func (r *report) writeTo(fn writeFunc) *report {
	r.write = fn
	return r
}

func (r *report) run() ([]result, error) {
	if r.retrieve == nil {
		panic("report retrieve cannot be nil")
	}

	results, err := r.retrieve(r.input)
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
	if r.write != nil {
		if err := r.write(r.output, results); err != nil {
			return nil, err
		}
	}

	return results, nil
}
