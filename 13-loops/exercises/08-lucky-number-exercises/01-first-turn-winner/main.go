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
"time"
"strconv"
"os"
"rand"
)

// ---------------------------------------------------------
// EXERCISE: First Turn Winner
//
//  If the player wins on the first turn, then display
//  a special bonus message.
//
// RESTRICTION
//  The first picked random number by the computer should
//  match the player's guess.
//
// EXAMPLE SCENARIO
//  1. The player guesses 6
//  2. The computer picks a random number and it happens
//     to be 6
//  3. Terminate the game and print the bonus message
// ---------------------------------------------------------

func main() {
  if len(os.Args)==1{
    fmt.Println("enter the guess number:")
    return
}
  b:=os.Args[1]
guess:=strconv.Atoi(b)
  rand.Seed(time.Now().UnixNano())
  n:=0
  for i:=0;i<5;i++{
    n=rand.Intn(guess+1)
    if n==guess{
      if i==0{
        fmt.Println("XXXXX---A special bonus message---XXXXXX")
       }else{
         fmt.Println(" XXXXX---LUCKY GAME WINNER---XXXXX")
       }
}
}
if n!=guess{
fmt.Println("BETTER LUCK NEXT TIME")
}



    
