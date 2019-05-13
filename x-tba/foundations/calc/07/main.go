package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// lesson: modulo operator and type conversion

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
	case "+", "plus":
		op, res = "+", a+b
	case "-", "minus":
		op, res = "-", a-b
	case "*", "times":
		op, res = "*", a*b
	case "/", "div":
		op, res = "/", a/b
	case "%", "mod":
		res = float64(int(a) % int(b))
	default:
		fmt.Println("Please provide a valid operation.")
		return
	}

	fmt.Printf("%v %v %v = %v\n", a, op, b, res)
}
