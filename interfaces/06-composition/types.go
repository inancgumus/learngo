package main

// don't separate your interfaces like this.
// put them in the same file where you need to use them.
// this is here for clarity

type printer interface {
	print()
}

// TODO: NEW
type summer interface {
	sum() money
}

// TODO: NEW
// interface embedding
// When an interface includes multiple methods,
// choose a name that accurately describes its purpose.
type item interface {
	printer
	summer
}
