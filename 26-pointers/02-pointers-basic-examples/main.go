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
	var (
		counter int
		V       int
		P       *int
	)

	counter = 100 // counter is an int variable
	P = &counter  // P is a pointer int variable
	V = *P        // V is a int variable (a copy of counter)

	if P == nil {
		fmt.Printf("P is nil and its address is %p\n", P)
	}

	if P == &counter {
		fmt.Printf("P is equal to counter's address: %p == %p\n", 
			P, &counter)
	}

	fmt.Printf("counter: %-13d addr: %-13p\n", counter, &counter)
	fmt.Printf("P      : %-13p addr: %-13p *P: %-13d\n", P, &P, *P)
	fmt.Printf("V      : %-13d addr: %-13p\n", V, &V)

	fmt.Println("\n••••• change counter")
	counter = 10 // V doesn't change because it's a copy
	fmt.Printf("counter: %-13d addr: %-13p\n", counter, &counter)
	fmt.Printf("V      : %-13d addr: %-13p\n", V, &V)

	fmt.Println("\n••••• change counter in passVal()")
	counter = passVal(counter)
	fmt.Printf("counter: %-13d addr: %-13p\n", counter, &counter)

	fmt.Println("\n••••• change counter in passPtrVal()")
	passPtrVal(&counter) // same as passPtrVal(&counter) (no need to return)
	passPtrVal(&counter) // same as passPtrVal(&counter) (no need to return)
	fmt.Printf("counter: %-13d addr: %-13p\n", counter, &counter)
}

// *pn is a int pointer variable (copy of P)
func passPtrVal(pn *int) {
	fmt.Printf("pn     : %-13p addr: %-13p *pn: %d\n", pn, &pn, *pn)

	// pointers can breach function isolation borders
	*pn++ // counter changes because `pn` points to `counter` — (*pn)++
	fmt.Printf("pn     : %-13p addr: %-13p *pn: %d\n", pn, &pn, *pn)

	// setting it to nil doesn't effect the caller (the main function)
	pn = nil
}

// n is a int variable (copy of counter)
func passVal(n int) int {
	n = 50 // counter doesn't change because `n` is a copy
	fmt.Printf("n      : %-13d addr: %-13p\n", n, &n)
	return n
}
