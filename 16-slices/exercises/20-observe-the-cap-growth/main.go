package main

// ---------------------------------------------------------
// EXERCISE: Observe the capacity growth
//
//  Write a program that loops 10 million times to append an element
//  to a slice, on each step of the loop. Observe the capacity.
//
//
// STEPS
//
//  1. Create a nil slice
//
//  2. Loop 10e6 times
//
//  3. On each loop step: Append an element to the slice
//
//  4. Only print the length and capacity of the slice everytime
//     the capacity changes.
//
//  5. Print also the growth rate by calculating the previous and
//     the current capacity.
//
//
// EXPECTED OUTPUT
//
//  len:0               cap:0               growth:NaN
//  len:1               cap:1               growth:+Inf
//  len:2               cap:2               growth:2.00
//  ... and so on.
//
// ---------------------------------------------------------

func main() {}
