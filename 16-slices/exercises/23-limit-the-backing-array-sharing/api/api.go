package api

// note: "client" means the users or importers of your package.

// The original temperatures slice.
var temps = []int{5, 10, 3, 25, 45, 80, 90}

//                ^  ^   ^  ^           ^
//                |  |   |  +-----------+
// the client is allowed to |
// change these elements.   |
//                          |
// but the client shouldn't change the
// rest of the elements after the 3rd element.

// Read returns a range of temperature readings beginning from
// the `start` until to the `stop`.
func Read(start, stop int) []int {
	// ----------------------------------------
	// RESTRICTIONS — ONLY ADD YOUR CODE HERE

	// returns a slice from the temps slice to the client.
	portion := temps[start:stop]

	// ----------------------------------------
	return portion
}

// All returns all the temperature readings
func All() []int {
	return temps
}
