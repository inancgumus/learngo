package main

import (
	"fmt"
	"os"
	"strings"
)

const yearly = 4

func main() {
	var books [yearly]string

	books[0] = "Kafka's Revenge"
	books[1] = "Stay Golden"
	books[2] = "Everythingship"
	books[3] += books[0] + " 2nd Edition"

	if len(os.Args) != 2 {
		fmt.Println("Tell me a book title")
		return
	}
	query := os.Args[1]

	fmt.Println("Search Results:")
	fmt.Println("---------------")

	var found bool

	for _, v := range books {
		if strings.Contains(strings.ToLower(v), query) {
			fmt.Println("+", v)
			found = true
		}
	}

	if !found {
		fmt.Printf("We don't have that book: %q\n", query)
	}
}
