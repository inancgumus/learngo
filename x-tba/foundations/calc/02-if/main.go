package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// lesson: if

	if len(os.Args) != 3 {
		fmt.Println("Usage: calc <number1> <number2>")
		return
	}

	a, _ := strconv.Atoi(os.Args[1])
	b, _ := strconv.Atoi(os.Args[2])

	fmt.Printf("%v + %v = %v\n", a, b, a+b)
}
