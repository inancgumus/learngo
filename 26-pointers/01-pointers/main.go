// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

func main() {
	var counter byte = 100
	P := &counter
	V := *P

	fmt.Printf("counter : %-16d address: %-16p\n", counter, &counter)
	fmt.Printf("P       : %-16p address: %-16p *P: %-16d\n", P, &P, *P)
	fmt.Printf("V       : %-16d address: %-16p\n", V, &V)

	V = 200
	fmt.Println()
	fmt.Printf("counter : %-16d address: %-16p\n", counter, &counter)
	fmt.Printf("P       : %-16p address: %-16p *P: %-16d\n", P, &P, *P)
	fmt.Printf("V       : %-16d address: %-16p\n", V, &V)

	V = counter // reset the V to counter's initial value
	counter++
	fmt.Println()
	fmt.Printf("counter : %-16d address: %-16p\n", counter, &counter)
	fmt.Printf("P       : %-16p address: %-16p *P: %-16d\n", P, &P, *P)
	fmt.Printf("V       : %-16d address: %-16p\n", V, &V)

	*P = 25
	fmt.Println()
	fmt.Printf("counter : %-16d address: %-16p\n", counter, &counter)
	fmt.Printf("P       : %-16p address: %-16p *P: %-16d\n", P, &P, *P)
	fmt.Printf("V       : %-16d address: %-16p\n", V, &V)
}
