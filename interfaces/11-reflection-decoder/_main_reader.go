// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	dec := decoder{
		"book":   book{},
		"game":   game{},
		"puzzle": puzzle{},
		"toy":    toy{},
	}

	store, err := fromReader(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print(store)
}
