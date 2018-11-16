package main

import "fmt"

// STORY:
// Hipster's Love store publishes limited books
// twice a year.
//
// The number of books they publish is fixed at 4.

// So, let's create a 4 elements string array for the books.

const (
	winter = 1
	summer = 3
	yearly = winter + summer
)

func main() {
	var (
		prevBooks [yearly]string
		books     [yearly]string
	)

	books[0] = "Kafka's Revenge"
	books[1] = "Stay Golden"
	books[2] = "Everythingship"
	books[3] += books[0] + " 2nd Edition"

	prevBooks = books
	books[0] = "Silver Ages"
	books[1] = "Dragon's Fire"
	books[2] = "Nothingless"
	books[3] = prevBooks[2] + " 2nd Edition"

	fmt.Println("Last Year's Books (Sold Out!):")
	for _, v := range prevBooks {
		fmt.Println("+", v)
	}

	fmt.Println("\nNew Books:")
	for _, v := range books {
		fmt.Println("+", v)
	}
}
