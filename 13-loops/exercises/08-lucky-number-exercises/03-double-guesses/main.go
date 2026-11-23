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
  "os"
  "math/rand"
  "strconv"
  "time"
  )

// ---------------------------------------------------------
// EXERCISE: Double Guesses
//
//  Let the player guess two numbers instead of one.
//
// HINT:
//  Generate random numbers using the greatest number
//  among the guessed numbers.
//
// EXAMPLES
//  go run main.go 5 6
//  Player wins if the random number is either 5 or 6.
// ---------------------------------------------------------

func main() {
  fmt.Println("-------GAME-----------GAME-------------GAME-------")
	if len(os.Args) == 1 {
		fmt.Println("Enter the guess number please:")
		return
	}
	a := os.Args[1:]
  max  := -2147483648
  if a[0]>a[1]{
    b :=a[0]
    }else{
    b :=a[1]
    }
    
  
	guess, err := strconv.Atoi(b)
	if err != nil {
		fmt.Println("xxx---ERROR---XXX")
		return
	}
	i := 0
	rand.Seed(time.Now().UnixNano())
	for n := 0; n < 5; n++ {
		i = rand.Intn(guess + 1)
		if i == guess {
			fmt.Println(guess)
			fmt.Println("XXXXXXXX---LUCKY NUMBER GENERATOR---XXXXXXXXX")
			break
		}
		fmt.Printf("%d ", i)
	}
	if i != guess {
		fmt.Println()
		fmt.Println("XXXXXXX---BETTER LUCK NEXT TIME---XXXXXXX")
	}
	fmt.Println()
}
