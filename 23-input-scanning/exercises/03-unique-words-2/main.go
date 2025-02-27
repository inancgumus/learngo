// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Unique Words 2
//
//  Use your solution from the previous "Unique Words"
//  exercise.
//
//  Before adding the words to your map, remove the
//  punctuation characters and numbers from them.
//
//
// BE CAREFUL
//
//  Now the shakespeare.txt contains upper and lower
//  case letters too.
//
//
// EXPECTED OUTPUT
//
//  go run main.go < shakespeare.txt
//
//   There are 100 words, 69 of them are unique.
//
// ---------------------------------------------------------

func main() {
	// This is the regular expression pattern you need to use:
	// [^A-Za-z]+
	//
	// Matches to any character but upper case and lower case letters
	total, words := 0, make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	re:=regexp.MustCompile('[^a-zA-Z\\s]+')
	for input.Scan() {
		text:=input.Text()
		text=strings.ToLower(text)
		text=re.ReplaceAllString(text,"")
		total++
		words[text]++
	}
	unique := len(words)
	fmt.Printf("There are %d , %d are them unique", total, unique)

}
