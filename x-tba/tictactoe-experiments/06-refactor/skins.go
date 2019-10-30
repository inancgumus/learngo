// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "github.com/fatih/color"

// skin options :-)
type skin struct {
	empty, mark1, mark2               string
	header, middle, footer, separator string

	unicode bool
}

var skins = map[string]skin{
	"colorful": colorfulSkin,
	"poseidon": poseidonSkin,
	"statues":  statuesSkin,
	"aliens":   aliensSkin,
	"snow":     snowSkin,
	"monkeys":  monkeysSkin,
}

var defaultSkin = skin{
	empty: "   ",
	mark1: " X ",
	mark2: " O ",

	header:    "┌───┬───┬───┐",
	middle:    "├───┼───┼───┤",
	footer:    "└───┴───┴───┘",
	separator: "│",
}

var colorfulSkin = skin{
	empty:     "   ",
	mark1:     color.CyanString(" X "),
	mark2:     color.HiMagentaString(" O "),
	header:    color.HiBlueString("┌───┬───┬───┐"),
	middle:    color.HiBlueString("├───┼───┼───┤"),
	footer:    color.HiBlueString("└───┴───┴───┘"),
	separator: color.BlueString("│"),
}

var poseidonSkin = skin{
	empty:     "❓  ",
	mark1:     "🔱  ",
	mark2:     "⚓️  ",
	header:    "●————●————●————●",
	middle:    "●————●————●————●",
	footer:    "●————●————●————●",
	separator: "⎮ ",
}

var statuesSkin = skin{
	empty:     "❓  ",
	mark1:     "🗿  ",
	mark2:     "🗽  ",
	header:    "┌────┬────┬────┐",
	middle:    "├────┼────┼────┤",
	footer:    "└────┴────┴────┘",
	separator: "│ ",
}

var aliensSkin = skin{
	empty:     "❓  ",
	mark1:     "👽  ",
	mark2:     "👾  ",
	header:    "┌────┬────┬────┐",
	middle:    "├────┼────┼────┤",
	footer:    "└────┴────┴────┘",
	separator: "│ ",
}

var snowSkin = skin{
	empty:     "❓  ",
	mark1:     "⛄ ️ ",
	mark2:     "❄️  ",
	header:    "╔════╦════╦════╗",
	middle:    "╠════╬════╬════╣",
	footer:    "╚════╩════╩════╝",
	separator: "║ ",
}

var monkeysSkin = skin{
	empty:     "🍌  ",
	mark1:     "🙈  ",
	mark2:     "🙉  ",
	header:    "┌────┬────┬────┐",
	middle:    "├────┼────┼────┤",
	footer:    "└────┴────┴────┘",
	separator: "│ ",
}
