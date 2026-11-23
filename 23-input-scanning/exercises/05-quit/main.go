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
  "strings"
  "os"
  )
// ---------------------------------------------------------
// EXERCISE: Quit
//
//  Create a program that quits when a user types the
//  same word twice.
//
//
// RESTRICTION
//
//  The program should work case insensitive.
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//
//   hey
//   HEY
//   TWICE!
//
//  go run main.go
//
//   hey
//   hi
//   HEY
//   TWICE!
// ---------------------------------------------------------

func main() {
  args:=os.Args[1:]
  
  if len(args)<2{
    fmt.Println("Errror")
    return
    }
  words:=make(map[string]int)
  for index,_:=range args{
    fmt.Println(args[index])
    key:=strings.ToLower(args[index])
    
    if _,ok:=words[key];ok{
      fmt.Println("TWICE!")
      return
      }
    }
}
