// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main
import(
  "fmt"
  "os"
  "math/rand"
  "strconv"
  "time"
  )

// ---------------------------------------------------------
// EXERCISE: Enough Picks
//
//  If the player's guess number is below 10;
//  then make sure that the computer generates a random
//  number between 0 and 10.
//
//  However, if the guess number is above 10; then the
//  computer should generate the numbers
//  between 0 and the guess number.
//
// WHY?
//  This way the game will be harder.
//
//  Because, in the current version of the game, if
//  the player's guess number is for example 3; then the
//  computer picks a random number between 0 and 3.
//
//  So, currently a player can easily win the game.
//
// EXAMPLE
//  Suppose that the player runs the game like this:
//    go run main.go 9
//
//  Or like this:
//    go run main.go 2
//
//    Then the computer should pick a random number
//    between 0-10.
//
//  Or, if the player runs it like this:
//    go run main.go 15
//
//    Then the computer should pick a random number
//    between 0-15.
// ---------------------------------------------------------

func main() {
  fmt.Println("-------GAME-----------GAME-------------GAME-------")
	if len(os.Args) == 1 {
		fmt.Println("Enter the guess number please:")
		return
	}
	a := os.Args[1]
  if a<10{
    a=10
    }
	guess, err := strconv.Atoi(a)
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
