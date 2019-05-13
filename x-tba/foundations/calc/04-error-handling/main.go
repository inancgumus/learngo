package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// lesson: error handling + short decl. assignment

	if len(os.Args) != 3 {
		fmt.Println("Usage: calc <number1> <number2>")
		return
	}

	a, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Please provide a valid number")
		return
	}

	b, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Println("Please provide a valid number")
		return
	}

	fmt.Printf("%v + %v = %v\n", a, b, a+b)
}
