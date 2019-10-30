// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// STORY:
// Hipster's Love store publishes limited books
// twice a year.
//
// The number of books they publish is fixed at 4.

// So, let's create a 4 elements string array for the books.

func main() {
	// Use this only when you don't know about the elements beforehand
	{
		var books [4]string

		books[0] = "Kafka's Revenge"
		books[1] = "Stay Golden"
		books[2] = "Everythingship"
		books[3] += "Kafka's Revenge 2nd Edition"

		_ = books
	}

	// This is not necessary, use the short declaration syntax below
	{
		var books = [4]string{
			"Kafka's Revenge",
			"Stay Golden",
			"Everythingship",
			"Kafka's Revenge 2nd Edition",
		}

		_ = books
	}

	// Use this if you know about the elements
	{
		books := [4]string{
			"Kafka's Revenge",
			"Stay Golden",
			"Everythingship",
			"Kafka's Revenge 2nd Edition",
		}

		_ = books
	}

	{
		// Use this if you know about the elements
		books := [4]string{
			"Kafka's Revenge",
			"Stay Golden",
		}

		// Uninitialized elements will be set to their zero values
		fmt.Printf("books  : %#v\n", books)
	}

	// You can also use the ellipsis syntax
	// ... equals to 4
	{
		books := [...]string{
			"Kafka's Revenge",
			"Stay Golden",
			"Everythingship",
			"Kafka's Revenge 2nd Edition",
		}

		_ = books
	}
}
