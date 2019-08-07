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
		group:  noopGrouper,
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
	//            input      filterBy    groupBy
	//           scanner  (result) bool  map[string]result
	//
	// stdin -> []result -> []results -> []result -> output(stdout)

	res, err := r.input()
	if err != nil {
		return nil, err
	}

	res = filterBy(res, r.filter)
	res = groupBy(res, r.group)
	err = r.output(res)

	return res, err
}
