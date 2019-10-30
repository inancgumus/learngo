// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	c "github.com/fatih/color"
	rainbow "github.com/guineveresaenger/golang-rainbow"
)

// TODO: move the main logic into a helper

func main() {
	const banner = `
~~~~~~~~~~~~~~~~
  TIC~TAC~TOE
~~~~~~~~~~~~~~~~`

	in := bufio.NewScanner(os.Stdin)
	sk := selectSkin(in)
	lg := logger{
		print:   fmt.Print,
		printf:  fmt.Printf,
		println: fmt.Println,
	}

	for {
		g := newGame(sk, lg)

	game:
		for {
			rainbow.Rainbow(banner, 3)
			g.print()
			fmt.Printf(c.CyanString("\n%s [1-9]: ", g.player))

			if !in.Scan() {
				break
			}

			pos, _ := strconv.Atoi(in.Text())

			switch st := g.play(pos); st {
			case stateAlreadyPlayed, stateWrongPosition:
				announce(g, st)
				continue
			case stateWon, stateTie:
				announce(g, st)
				break game
			}
			g.Print("\033[2J")

		}

		fmt.Print(c.MagentaString("One more game? [y/n]: "))
		if in.Scan(); in.Text() != "y" {
			fmt.Println("OK, bye!")
			break
		}
	}
}

func announce(g *game, st state) {
	red := c.New(c.FgRed, c.Bold)
	green := c.New(c.BgBlack, c.FgGreen, c.Bold)

	switch st {
	case stateAlreadyPlayed, stateWrongPosition:
		red.Printf("\n>>> You can't play there!\n")
	case stateWon:
		g.print()
		green.Printf("\nWINNER: %s\n", g.player)
	case stateTie:
		g.print()
		green.Printf("\nTIE!\n")
	}
}

func selectSkin(in *bufio.Scanner) skin {
	fmt.Println(c.MagentaString("Our finest selection of skins:"))
	for name := range skins {
		fmt.Printf("- %s\n", name)
	}

	fmt.Print(c.GreenString("\nEnter the name of the skin: "))

	in.Scan()
	if sk, ok := skins[in.Text()]; ok {
		return sk
	}
	return defaultSkin
}
