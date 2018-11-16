package main

import "fmt"

// STORY:
// You want to compare two bookcases,
// whether they're equal or not.

func main() {
	// When comparing two arrays,
	// Their types should be identical

	var (
		// equal (types + elements are identical)::
		blue = [3]int{6, 9, 3}
		red  = [3]int{6, 9, 3}

		// equal (types + elements are identical):
		// blue = [...]int{6, 9, 3}
		// red  = [3]int{6, 9, 3}

		// not equal (element ordering are different):
		// blue = [3]int{6, 9, 3}
		// red  = [3]int{3, 9, 6}

		// not equal (the last elements are not equal):
		// blue = [3]int{6, 9}
		// red  = [3]int{6, 9, 3}

		// not comparable (type mismatch: length):
		// blue = [3]int{6, 9, 3}
		// red  = [5]int{6, 9, 3}

		// not comparable (type mismatch: element type):
		// blue = [3]int64{6, 9, 3}
		// red  = [3]int{6, 9, 3}
	)
	fmt.Printf("blue bookcase : %v\n", blue)
	fmt.Printf("red bookcase  : %v\n", red)

	fmt.Println("Are they equal?", blue == red)
}
