// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

func main() {
	// removed error handling for brevity (normally you shouldn't do this).

	// encode to memory (to a byte slice)
	// out, _ := encode()
	// fmt.Println(string(out))

	// read from memory (from a byte slice):
	// l, _ := decode(out)

	// read from a file by name:
	// l, _ := decodeFile("database.json")

	// read from a file #1 (not good):
	// f, _ := os.Open("database.json")
	// l, _ := decodeFileObject(f)
	// f.Close()

	// read from a file #2 (better):
	// f, _ := os.Open("database.json")
	// l, _ := decodeReader(f)
	// f.Close()

	// read from memory (from a string):
	// const data = `[
	// { "title": "moby dick", "price": 10, "released": 118281600 },
	// { "title": "odyssey", "price": 15, "released": 733622400 },
	// { "title": "hobbit", "price": 25, "released": -62135596800 }
	// ]`
	// l, _ := decodeReader(strings.NewReader(data))

	// read from inanc's web server:
	// res, _ := http.Get("https://inancgumus.github.io/x/database.json")
	// l, _ := decodeReader(res.Body)
	// res.Body.Close()

	// fmt.Print(l)
}
