package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// lesson: switch

	if len(os.Args) != 4 {
		fmt.Println("Usage: calc <number1> <operator> <number2>")
		return
	}

	a, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Please provide a valid number")
		return
	}

	b, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Println("Please provide a valid number")
		return
	}

	var (
		op  = os.Args[2]
		res float64
	)

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	}

	fmt.Printf("%v %v %v = %v\n", a, op, b, res)
}
