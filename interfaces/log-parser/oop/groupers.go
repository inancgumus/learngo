package main

type groupFunc func(record) string

// domainGrouper groups by domain.
// but it keeps the other fields.
// for example: it returns pages as well, but you shouldn't use them.
// exercise: write a function that erases the unnecessary data.
func domainGrouper(r record) string {
	return r.domain
}

func pageGrouper(r record) string {
	return r.domain + r.page
}
