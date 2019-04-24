// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	var (
		counter int
		V       int
		P       *int
	)

	if P == nil {
		fmt.Printf("P is nil and its address is %p\n", P)
	}

	counter = 100 // counter is an int variable
	P = &counter  // P is a pointer int variable
	V = *P        // V is a int variable (a copy of counter)

	if P == &counter {
		fmt.Printf("P's address is equal to counter: %p == %p\n", P, &counter)
	}

	fmt.Printf("counter : %-16d address: %-16p\n", counter, &counter)
	fmt.Printf("P       : %-16p address: %-16p *P: %-16d\n", P, &P, *P)
	fmt.Printf("V       : %-16d address: %-16p\n", V, &V)

	fmt.Println("\n••••• CHANGE: counter")
	counter = 10 // V doesn't change because it's a copy
	fmt.Printf("counter : %-16d address: %-16p\n", counter, &counter)
	fmt.Printf("V       : %-16d address: %-16p\n", V, &V)

	fmt.Println("\n••••• CHANGE IN: passVal")
	counter = passVal(counter)
	fmt.Printf("counter : %-16d address: %-16p\n", counter, &counter)

	fmt.Println("\n••••• CHANGE IN: passPtrVal")
	passPtrVal(&counter) // same as passPtrVal(&counter) (no need to return)
	fmt.Printf("counter : %-16d address: %-16p\n", counter, &counter)
}

// *pn is a int pointer variable (copy of P)
func passPtrVal(pn *int) {
	fmt.Printf("pn      : %-16p address: %-16p *pn: %d\n", pn, &pn, *pn)

	// pointers can breach function isolation borders
	*pn++ // counter changes because `pn` points to `counter` — (*pn)++
	fmt.Printf("pn      : %-16p address: %-16p *pn: %d\n", pn, &pn, *pn)
}

// n is a int variable (copy of counter)
func passVal(n int) int {
	n = 50 // counter doesn't change because `n` is a copy
	fmt.Printf("n       : %-16d address: %-16p\n", n, &n)
	return n
}
