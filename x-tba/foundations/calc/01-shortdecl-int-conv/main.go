package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// lesson: multi-return funcs, %v, and _

	a, _ := strconv.Atoi(os.Args[1])
	b, _ := strconv.Atoi(os.Args[2])

	fmt.Printf("%v + %v = %v\n", a, b, a+b)
}
