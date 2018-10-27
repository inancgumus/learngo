// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"strings"
)

func main() {
	msg := `
	
	The weather looks good.
I should go and play.



	`

	fmt.Println(strings.TrimSpace(msg))
}
