package main

import "fmt"

func main() {
	//  1. Declare and printf the following arrays:
	//   1. A string array with 4 items
	//   2. An int array 5 items
	//   3. A byte array with 5 items
	//   4. A float64 array with 1 item
	//   5. A bool array with 4 items
	//   6. A byte array without any items
	var (
		names     [4]string
		distances [5]int
		data      [5]byte
		ratios    [1]float64
		switches  [4]bool
		zero      [0]bool
	)

	fmt.Printf("names    : %#v\n", names)
	fmt.Printf("distances: %#v\n", distances)
	fmt.Printf("data     : %#v\n", data)
	fmt.Printf("ratios   : %#v\n", ratios)
	fmt.Printf("switches : %#v\n", switches)
	fmt.Printf("zero     : %#v\n", zero)

	//  2. Print the types of the previous arrays.
	fmt.Printf("names    : %T\n", names)
	fmt.Printf("distances: %T\n", distances)
	fmt.Printf("data     : %T\n", data)
	fmt.Printf("ratios   : %T\n", ratios)
	fmt.Printf("switches : %T\n", switches)
	fmt.Printf("zero     : %T\n", zero)
}
