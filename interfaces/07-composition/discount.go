package main

// usually: include the discount method in the list.go
// it is here for clarity

// TODO: NEW
func (l list) discount(ratio float64) {
	type discounter interface {
		discount(float64)
	}

	for _, it := range l {
		if it, ok := it.(discounter); ok {
			it.discount(ratio)
		}
	}
}
