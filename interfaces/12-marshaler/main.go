// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

const data = `[
 {
  "title": "moby dick",
  "price": 10,
  "released": 118281600
 },
 {
  "title": "odyssey",
  "price": 15,
  "released": 733622400
 },
 {
  "title": "hobbit",
  "price": 25,
  "released": -62135596800
 }
]`

func main() {
	/* encoding */
	l := list{
		{Title: "moby dick", Price: 10, Released: toTimestamp(118281600)},
		{Title: "odyssey", Price: 15, Released: toTimestamp("733622400")},
		{Title: "hobbit", Price: 25},
	}

	data, err := json.MarshalIndent(l, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

	/* decoding */
	// var l list
	// if err := json.Unmarshal([]byte(data), &l); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Print(l)
}

/*
Summary:

- json.Marshal() and json.MarshalIndent() can only encode primitive types.
  - Custom types can tell the encoder how to encode.
  - To do that satisfy the json.Marshaler interface.

- json.Unmarshal() can only decode primitive types.
  - Custom types can tell the decoder how to decode.
  - To do that satisfy the json.Unmarshaler interface.

- strconv.AppendInt() can append an int value to a []byte.
  - There are several other functions in the strconv package for other primitive types as well.
  - Do not make unnecessary string <-> []byte conversions.

- log.Fatal() can print the given error message and terminate the program.
  - Use it only from the main(). Do not use it in other functions.
  - main() is should be the main driver of a program.
*/
