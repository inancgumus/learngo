package main

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Assign the Arrays
//
//  1. Create an array named books
//  2. Add book titles to the array
//  3. Create two more copies of the array named: upper and lower
//  4. Change the book titles to uppercase in the upper array only
//  5. Change the book titles to lowercase in the lower array only
//  6. Print all the arrays
//
// NOTE
//  Check out the strings package, it has function to convert cases to
//  upper and lower cases.
//
// BONUS
//  1. Invent your own arrays with different types other than string,
//     and do some manipulations on them.
//
//  👉 THISSSS--------------------------------------------------------
//  2. Find some Turkish book titles and do the same upper, lowercase conversion
//     for them.
//
//     Here are some books: https://www.goodreads.com/group/bookshelf/417154-bilimkurgu-kul-b?shelf=read
//
//     Note: You'd need to use special functions to convert the Turkish letters.
//           They're in the strings package as well.
//
// EXPECTED OUTPUT
//  ?
// ---------------------------------------------------------

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
