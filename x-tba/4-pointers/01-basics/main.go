// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
)

func main() {
	// As you've learned before Go is a pass-by-value programming language. So, when you assign a variable to another one, or pass it to a function, it gets copied.

	// #1: Let's take a look at an example.
	// Let's say you have an int variable.
	// You want a function to increase the number.
	number := 1

	// In that case, you need to use a pointer because a pointer can send the memory location of the variable to the function, so the function can change the value directly on the memory.
	//
	// For example, let me pass the memory address of the number to the function like so.
	//
	// Remember: Ampersand sign here returns the memory location of the given variable.
	incr(&number)

	// Let me print the variable.
	fmt.Println("&main.number     :", number)

	// #1C: Let's check it out.
	// As you can see, now the number has changed to 2.
	// The function could increment the number because it can directly access to the value of the variable by using its memory address.

	// If you look at the memory address of the number variable, you will see that the function sees the same address, let me show you.
	//
	// Let me print the address of the number variable like so.
	// Every pointer stores a memory location value.
	// So the %p verb prints the address that a pointer stores.
	fmt.Printf("&main.number     : %p\n", &number)
}

// #1B: Let me also declare the function like so.
// This time, it accepts an int pointer.
// When you put an asterisk sign in front of a type, it becomes a pointer.
// Every type comes with its own pointer type like so.
// This is an int pointer because it has an asterisk in front of the int type. So, this pointer can only store the memory address of an int value.
func incr(num *int) {
	// Inside the function I'm going to change the num variable by deferencing it like so, then I'm going to increase it as usual.
	*num++

	// #1D: Let me also print the memory address that the num pointer stores.
	// Remember: num here is a pointer variable, it stores a memory address of an int value.
	fmt.Printf("incr.number      : %p\n", num)

	// As you can see, the main function and the incr function use the same variable. They both are looking at the same memory location. That's why this function can change the value.

	// However, if I print the memory address of the local num variable, it will print a different address, let me show you.
	fmt.Printf("&incr.number     : %p\n", &num)

	// As you can see, the memory address of the local num variable is different than the memory address of this variable. It's because, Go is a pass-by-value language. So, it even copies the pointer variables. So, the local num pointer variable is a new variable inside this function. However, it stores the same address of the number variable of main function.

	// In Go, everything is passed by value, even pointers, however, when the functions knows the memory address of a variable, it can change it directly. Otherwise it cannot change it because the variable belongs to the local scope of the function.
}
