package main

type groupFunc func(result) string

// domainGrouper groups by domain.
// but it keeps the other fields.
// for example: it returns pages as well, but you shouldn't use them.
// exercise: write a function that erases the unnecessary data.
func domainGrouper(r result) string {
	return r.domain
}

func pageGrouper(r result) string {
	return r.domain + r.page
}
