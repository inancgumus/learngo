// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Refactor the game store to funcs - step #2
//
//  Let's continue the refactoring from the previous
//  exercise. This time, you're going to refactor the
//  command handling logic.
//
//
//  Create commands.go file
//
//     1- Add a func that runs the given command from the user:
//
//        Name  : runCmd
//        Inputs: input string, []game, map[int]game
//        Output: bool
//
//        This func returns true if it wants the program to
//        continue. When it returns false, the program will
//        terminate. So, all the commands that it calls need
//        to return true or false as well depending on whether
//        they want to cause the program to terminate or not.
//
//     2- Add a func that handles the quit command:
//
//        Name  : cmdQuit
//        Input : none
//        Output: bool
//
//     3- Add a func that handles the list command:
//
//        Name  : cmdList
//        Inputs: []game
//        Output: bool
//
//     4- Add a func that handles the id command:
//
//        Name  : cmdByID
//        Inputs: cmd []string, []game, map[int]game
//        Output: bool
//
//     5- Refactor the runCmd() with the cmdXXX funcs.
//
//  Go back to main.go and change the existing code with
//  the new funcs that you've created. There are hints
//  inside the code.
//
// ---------------------------------------------------------

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	games := load()
	byID := indexByID(games)

	fmt.Printf("Inanc's game store has %d games.\n", len(games))

	in := bufio.NewScanner(os.Stdin)
	for {
		// menu()
		fmt.Printf(`
  > list   : lists all the games
  > id N   : queries a game by id
  > quit   : quits

`)

		if !in.Scan() {
			break
		}

		// --- runCmd start ---
		fmt.Println()

		cmd := strings.Fields(in.Text())
		if len(cmd) == 0 {
			continue
		}

		switch cmd[0] {
		case "quit":
			// cmdQuit()
			fmt.Println("bye!")
			return

		case "list":
			// cmdList()
			for _, g := range games {
				printGame(g)
			}

		case "id":
			// cmdByID
			if len(cmd) != 2 {
				fmt.Println("wrong id")
				continue
			}

			id, err := strconv.Atoi(cmd[1])
			if err != nil {
				fmt.Println("wrong id")
				continue
			}

			g, ok := byID[id]
			if !ok {
				fmt.Println("sorry. I don't have the game")
				continue
			}

			printGame(g)
		}
		// --- runCmd end ---
	}
}
