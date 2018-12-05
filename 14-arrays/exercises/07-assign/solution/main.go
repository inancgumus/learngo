package main

import (
	"fmt"
	"strings"
)

func main() {
	books := [...]string{"Kafka's Revenge", "Stay Golden", "Everythingship"}

	upper, lower := books, books

	for i := range books {
		upper[i] = strings.ToUpper(upper[i])
		lower[i] = strings.ToLower(lower[i])
	}

	fmt.Printf("books: %q\n", books)
	fmt.Printf("upper: %q\n", upper)
	fmt.Printf("lower: %q\n", lower)
}
