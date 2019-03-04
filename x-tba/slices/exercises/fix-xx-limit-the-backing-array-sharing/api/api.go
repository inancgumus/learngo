package api

var temps = []int{5, 10, 3, 25, 45, 80, 90}

// Read returns a range of temperature readings beginning from
// the `start` until to the `stop`.
func Read(start, stop int) []int {
	// ----------------------------------------
	// RESTRICTIONS — ONLY ADD YOUR CODE HERE

	portion := temps[start:stop]

	// ----------------------------------------
	return portion
}

// All returns all the temperature readings
func All() []int {
	return temps
}
