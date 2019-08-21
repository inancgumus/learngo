package main

const fieldsLength = 4

type result struct {
	domain  string
	page    string
	visits  int
	uniques int
}

func (r result) add(other result) result {
	r.visits += other.visits
	r.uniques += other.uniques
	return r
}
