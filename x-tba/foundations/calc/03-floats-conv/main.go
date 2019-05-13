package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// lesson: floats encompass integers too

	if len(os.Args) != 3 {
		fmt.Println("Usage: calc <number1> <number2>")
		return
	}

	a, _ := strconv.ParseFloat(os.Args[1], 64)
	b, _ := strconv.ParseFloat(os.Args[2], 64)

	fmt.Printf("%v + %v = %v\n", a, b, a+b)
}
