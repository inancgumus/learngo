package api

var temps = []int{5, 10, 3, 25, 45, 80, 90}

// Read returns a range of temperature readings beginning from
// the `start` until to the `stop`.
func Read(start, stop int) []int {
	//
	// Uses a full slice expression to control the length of the
	// backing array (or the capacity of the returned slice).
	//
	// So the next append allocates a new backing array, which
	// in turn doesn't overwrite the temps slice's backing array.
	//                           ^^
	//                           ||
	//                          /  \
	//                         |    |
	portion := temps[start:stop:stop]
	return portion
}

// All returns all the temperature readings
func All() []int {
	return temps
}
