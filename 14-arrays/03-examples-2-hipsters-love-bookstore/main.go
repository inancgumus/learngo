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
	var books [yearly]string

	books[0] = "Kafka's Revenge"
	books[1] = "Stay Golden"
	books[2] = "Everythingship"
	books[3] += books[0] + " 2nd Edition"
	fmt.Printf("books  : %#v\n", books)

	// ------------------------------------
	// SEASONAL BOOKS
	// ------------------------------------

	var (
		wBooks [winter]string
		sBooks [summer]string
	)

	// assign the first book as the wBook's first book
	wBooks[0] = books[0]

	// assign all the books starting from the 2nd book
	// to sBooks array

	// sBooks[0] = books[1]
	// sBooks[1] = books[2]
	// sBooks[2] = books[3]

	// for i := 0; i < len(sBooks); i++ {
	// 	sBooks[i] = books[i+1]
	// }

	for i := range sBooks {
		sBooks[i] = books[i+1]
	}

	fmt.Printf("\nwinter : %#v\n", wBooks)
	fmt.Printf("\nsummer : %#v\n", sBooks)

	// ------------------------------------
	// Let's print the published books
	// ------------------------------------

	// equal to: [4]bool
	var published [len(books)]bool

	published[0] = true
	published[len(books)-1] = true

	fmt.Println("\nPublished Books:")
	for i, ok := range published {
		if ok {
			fmt.Printf("+ %s\n", books[i])
		}
	}
}
