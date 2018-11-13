// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import s "github.com/inancgumus/prettyslice"

func main() {
	ships := []string{
		"Normandy", "Verrikan", "Nexus",
	}

	// frigates := ships[:2]
	frigates := ships[:2:2]
	frigates = append(frigates, "Warsaw")

	s.Show("Ships", ships)
	s.Show("Frigates", frigates)
}
