package main

import "fmt"

func main() {
	var (
		myAge  byte
		herAge byte

		// uncomment the code below and observe the error
		// wrongDeclaration [-1]byte

		zero    [0]byte
		ages    [2]byte
		agesExp [2 + 1]byte

		tags [5]string
	)

	fmt.Printf("%d %d\n", myAge, herAge)
	fmt.Printf("%#v\n", zero)
	fmt.Printf("%#v\n", ages)
	fmt.Printf("%#v\n", agesExp)
	fmt.Printf("%#v\n", tags)
}
