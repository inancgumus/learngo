package main

// domainGrouper groups by domain.
// but it keeps the other fields.
// for example: it returns pages as well, but you shouldn't use them.
// exercise: write a function that erases superfluous data.
func domainGrouper(r result) string {
	return r.domain
}

func pageGrouper(r result) string {
	return r.domain + r.page
}

// groupBy allocates map unnecessarily
func noopGrouper(r result) string {
	// with something like:
	// return randomStrings()
	return ""
}
