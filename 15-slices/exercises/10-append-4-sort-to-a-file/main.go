package main

// ---------------------------------------------------------
// EXERCISE: Append #4 — Sort to a file
//
//  1. Get arguments from command-line
//
//  2. Sort them
//
//  3. Write the sorted slice to a file
//
//
// EXPECTED OUTPUT
//
//   go run main.go
//     Send me some items and I will sort them
//
//   go run main.go orange banana apple
//
//   cat sorted.txt
//     apple
//     banana
//     orange
//
//
// HINTS
//
//   ONLY READ THIS IF YOU GET STUCK
//
//   Below, []string means string slice, []byte means byte slice.
//
//   + You can use the os.Args[1:] to get a []string
//   + Then you can sort it using sort.Strings
//   + Use ioutil.WriteFile to write to a file.
//   + But you need to convert []string to []byte to be able to
//     write it to a file using the ioutil.WriteFile.
//   + To do that, create a new []byte and append the elements of your
//     []string.
// ---------------------------------------------------------

func main() {
}
