// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main
import (
  "fmt"
  "bufio"
  "os"
  "strings"
  )
// ---------------------------------------------------------
// EXERCISE: Uppercaser
//
//  Use a scanner to convert the lines to uppercase, and
//  print them.
//
//  1. Feed the shakespeare.txt to your program.
//
//  2. Scan the input using a new Scanner.
//
//  3. Print each line.
//
// EXPECTED OUTPUT
//  Please run the solution to see the expected output.
// ---------------------------------------------------------

func main() {
  in:=bufio.NewScanner(os.Stdin)
  for in.Scan(){
    fmt.Println(strings.ToUpper(in.Text()))
}
