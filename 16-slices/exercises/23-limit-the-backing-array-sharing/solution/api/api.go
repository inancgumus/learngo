package api

// The original temperatures slice.
var temps = []int{5, 10, 3, 25, 45, 80, 90}

// Read returns a range of temperature readings beginning from
// the `start` until to the `stop`.
func Read(start, stop int) []int {
	//
	// This third index prevents the clients of this package from
	// overwriting the original temps slice's backing array. It
	// limits the capacity of the returned slice. See the
	// full slice expressions lecture.
	//                            ^
	//                            |
	portion := temps[start:stop:stop]

	return portion
}

// All returns all the temperature readings
func All() []int {
	return temps
}
